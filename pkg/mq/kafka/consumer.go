package kafka

import (
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/zxq97/gokit/pkg/concurrent"
	"github.com/zxq97/gokit/pkg/mq"
)

var _ mq.Consumer = (*Consumer)(nil)

type Consumer struct {
	*mq.ConsumerMeta
	consumer *cluster.Consumer
	ch       chan *sarama.ConsumerMessage
}

func NewConsumer(conf *Config, topics []string, group string, fn mq.HandleFunc, opts ...mq.Option) (mq.Consumer, error) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.CommitInterval = time.Second
	if err := config.Validate(); err != nil {
		return nil, err
	}
	consumer, err := cluster.NewConsumer(conf.Addr, group, topics, config)
	if err != nil {
		return nil, err
	}
	meta := &mq.ConsumerMeta{Name: group, Fn: fn, NPoll: 1, NProc: 1, Done: make(chan struct{})}
	for _, o := range opts {
		o(meta)
	}
	return &Consumer{
		meta,
		consumer,
		make(chan *sarama.ConsumerMessage, mq.DefaultHandlerQueueSize),
	}, nil
}

func (c *Consumer) Start() error {
	rwg := sync.WaitGroup{}
	for i := 0; i < c.NPoll; i++ {
		rwg.Add(1)
		go func() {
			defer rwg.Done()
			c.readLoop()
		}()
	}
	go func() {
		rwg.Wait()
		close(c.ch)
	}()
	pwg := sync.WaitGroup{}
	for i := 0; i < c.NProc; i++ {
		pwg.Add(1)
		go func() {
			defer pwg.Done()
			c.procLoop()
		}()
	}
	go func() {
		pwg.Wait()
		close(c.Done)
	}()
	return nil
}

func (c *Consumer) Stop() error {
	if err := c.consumer.Close(); err != nil {
		return err
	}
	<-c.Done
	return c.consumer.CommitOffsets()
}

func (c *Consumer) readLoop() {
	for {
		select {
		case msg, ok := <-c.consumer.Messages():
			if ok {
				c.ch <- msg
			} else {
				return
			}
		case err, ok := <-c.consumer.Errors():
			if ok {
				mq.ExcLogger.Println("consumer read loop err", c.Name, err)
			}
		case ntf, ok := <-c.consumer.Notifications():
			if ok {
				mq.ExcLogger.Println("consumer read loop ntf", c.Name, ntf)
			}
		}
	}
}

func (c *Consumer) procLoop() {
	for {
		select {
		case msg, ok := <-c.ch:
			if ok {
				kafkaMsg, err := mq.Unmarshal(msg.Value)
				if err != nil {
					mq.ExcLogger.Println("consumer proc loop unmarshal err", c.Name, string(msg.Value), err)
					continue
				}
				ctx, cancel := mq.ConsumerContext(kafkaMsg.TraceId, c.ProcTimeout)
				now := time.Now()
				concurrent.F(func() {
					err = c.Fn(ctx, kafkaMsg)
					cancel()
				})
				if err != nil {
					mq.ExcLogger.Println("consumer proc loop fn err", c.Name, kafkaMsg, err)
					mq.ObserveMsg(msg.Topic, c.Name, "fail", err.Error(), now)
				} else {
					mq.ObserveMsg(msg.Topic, c.Name, "success", "", now)
				}
				c.consumer.MarkOffset(msg, "")
			} else {
				return
			}
		}
	}
}

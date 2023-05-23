package kafka

import (
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/zxq97/gokit/pkg/cast"
	"github.com/zxq97/gokit/pkg/concurrent"
	"github.com/zxq97/gokit/pkg/mq"
)

var _ mq.Consumer = (*Dispatch)(nil)

type Dispatch struct {
	nPoll    int
	fn       mq.HandleFunc
	consumer *cluster.Consumer
	chs      []chan *sarama.ConsumerMessage
	done     chan struct{}
}

func NewDispatchConsumer(conf *Config, topics []string, group string, fn mq.HandleFunc, n int) (*Dispatch, error) {
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
	chs := make([]chan *sarama.ConsumerMessage, n)
	for i := range chs {
		chs[i] = make(chan *sarama.ConsumerMessage, mq.DefaultHandlerQueueSize)
	}
	return &Dispatch{
		nPoll:    n,
		fn:       fn,
		consumer: consumer,
		chs:      chs,
		done:     make(chan struct{}),
	}, nil
}

func (c *Dispatch) Start() error {
	go c.readLoop()
	wg := sync.WaitGroup{}
	wg.Add(c.nPoll)
	for i := 0; i < c.nPoll; i++ {
		idx := i
		go func() {
			defer wg.Done()
			c.procLoop(idx)
		}()
	}
	go func() {
		wg.Wait()
		close(c.done)
	}()
	return nil
}

func (c *Dispatch) Stop() error {
	if err := c.consumer.Close(); err != nil {
		return err
	}
	<-c.done
	return c.consumer.CommitOffsets()
}

func (c *Dispatch) readLoop() {
	for {
		select {
		case msg, ok := <-c.consumer.Messages():
			if ok {
				idx := int(cast.ParseInt(string(msg.Key), 0) % int64(c.nPoll))
				c.chs[idx] <- msg
			} else {
				for i := 0; i < c.nPoll; i++ {
					close(c.chs[i])
				}
				return
			}
		case err, ok := <-c.consumer.Errors():
			if ok {
				mq.ExcLogger.Println("consumer read loop err", err)
			}
		case ntf, ok := <-c.consumer.Notifications():
			if ok {
				mq.ExcLogger.Println("consumer read loop ntf", ntf)
			}
		}
	}
}

func (c *Dispatch) procLoop(idx int) {
	for {
		select {
		case msg, ok := <-c.chs[idx]:
			if ok {
				kafkaMsg, err := mq.Unmarshal(msg.Value)
				if err != nil {
					mq.ExcLogger.Println("consumer proc loop unmarshal err", string(msg.Value), err)
					continue
				}
				ctx, cancel := mq.ConsumerContext(kafkaMsg.TraceId, time.Second)
				now := time.Now()
				concurrent.F(func() {
					err = c.fn(ctx, kafkaMsg)
					cancel()
				})
				if err != nil {
					mq.ExcLogger.Println("consumer proc loop fn err", kafkaMsg, err)
					mq.ObserveMsg(msg.Topic, "", "fail", err.Error(), now)
				} else {
					mq.ObserveMsg(msg.Topic, "", "success", "", now)
				}
				c.consumer.MarkOffset(msg, "")
			} else {
				return
			}
		}
	}
}

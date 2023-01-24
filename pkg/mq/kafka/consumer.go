package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/zxq97/gokit/pkg/concurrent"
)

type HandleFunc func(ctx context.Context, msg *KafkaMessage) error

type Consumer struct {
	name        string
	consumer    *cluster.Consumer
	fn          HandleFunc
	ch          chan *sarama.ConsumerMessage
	done        chan struct{}
	NPoll       int
	NProc       int
	procTimeout time.Duration
}

// NewConsumer
// f need process Idempotent
func NewConsumer(conf *Config, topics []string, group, name string, f HandleFunc, nPoll, nProc int, d time.Duration) (*Consumer, <-chan struct{}, error) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	consumer, err := cluster.NewConsumer(conf.Addr, group, topics, config)
	if err != nil {
		return nil, nil, err
	}
	done := make(chan struct{})
	return &Consumer{
		name:        name,
		consumer:    consumer,
		fn:          f,
		NPoll:       nPoll,
		NProc:       nProc,
		ch:          make(chan *sarama.ConsumerMessage, defaultHandlerQueueSize),
		done:        done,
		procTimeout: d,
	}, done, nil
}

func (c *Consumer) Start() {
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
		close(c.done)
	}()
}

func (c *Consumer) Stop() error {
	return c.consumer.Close()
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
				excLogger.Println("consumer read loop err", c.name, err)
			}
		case ntf, ok := <-c.consumer.Notifications():
			if ok {
				apiLogger.Println("consumer read loop ntf", c.name, ntf)
			}
		}
	}
}

func (c *Consumer) procLoop() {
	for {
		select {
		case msg, ok := <-c.ch:
			if ok {
				kafkaMsg, err := unmarshal(msg.Value)
				if err != nil {
					excLogger.Println("consumer proc loop unmarshal err", c.name, string(msg.Value), err)
					continue
				}
				ctx, cancel := consumerContext(kafkaMsg.TraceId, c.procTimeout)
				now := time.Now()
				concurrent.F(func() {
					err = c.fn(ctx, kafkaMsg)
					cancel()
				})
				if err != nil {
					excLogger.Println("consumer proc loop fn err", c.name, kafkaMsg, err)
					observeMsg(msg, c.name, "fail", err.Error(), now)
				} else {
					observeMsg(msg, c.name, "success", "", now)
				}
				c.consumer.MarkOffset(msg, "")
			} else {
				return
			}
		}
	}
}

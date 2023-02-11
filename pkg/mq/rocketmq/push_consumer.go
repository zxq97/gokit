package rocketmq

import (
	"context"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/concurrent"
	"github.com/zxq97/gokit/pkg/mq"
)

var _ mq.Consumer = (*PushConsumer)(nil)

type PushConsumer struct {
	*mq.ConsumerMeta
	consumer rocketmq.PushConsumer
}

func NewConsumer(conf *Config, topic, group string, fn mq.HandleFunc, opts ...mq.Option) (mq.Consumer, error) {
	meta := &mq.ConsumerMeta{Name: group, Fn: fn}
	for _, o := range opts {
		o(meta)
	}
	opt := []consumer.Option{
		// fixme consumer disorder consumer?
		consumer.WithConsumerOrder(true),
		consumer.WithNameServer(conf.NsAddr),
		consumer.WithGroupName(group),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
	}
	c, err := rocketmq.NewPushConsumer(opt...)
	if err != nil {
		return nil, err
	}
	f := func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, v := range ext {
			msg := &mq.MqMessage{}
			err := proto.Unmarshal(v.Body, msg)
			if err != nil {
				mq.ExcLogger.Println("consumer proc loop unmarshal err", v.Topic, string(v.Body), err)
				continue
			}
			ctx, cancel := mq.ConsumerContext(msg.TraceId, meta.ProcTimeout)
			now := time.Now()
			concurrent.F(func() {
				err = meta.Fn(ctx, msg)
				cancel()
			})
			if err != nil {
				mq.ExcLogger.Println("consumer proc loop fn err", msg, err)
				mq.ObserveMsg(v.Topic, v.MsgId, "fail", err.Error(), now)
			} else {
				mq.ObserveMsg(v.Topic, v.MsgId, "success", "", now)
			}
		}
		return consumer.ConsumeSuccess, nil
	}
	if err = c.Subscribe(topic, consumer.MessageSelector{Type: consumer.TAG}, f); err != nil {
		return nil, err
	}
	return &PushConsumer{
		meta,
		c,
	}, nil
}

func (c *PushConsumer) Start() error {
	return c.consumer.Start()
}

func (c *PushConsumer) Stop() error {
	return c.consumer.Shutdown()
}

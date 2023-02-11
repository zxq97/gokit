package rocketmq

import (
	"context"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/mq"
)

type Producer struct {
	producer rocketmq.Producer
}

func NewProducer(conf *Config, opts ...producer.Option) (*Producer, error) {
	opts = append(opts,
		producer.WithNameServer(conf.NsAddr),
		producer.WithSendMsgTimeout(time.Second),
	)
	p, err := rocketmq.NewProducer(opts...)
	if err != nil {
		return nil, err
	}
	return &Producer{producer: p}, p.Start()
}

func (p *Producer) SendMessage(ctx context.Context, topic, key, tag string, msg proto.Message, delayLvl int) error {
	bs, err := mq.WarpMessage(ctx, tag, msg)
	pMsg := primitive.NewMessage(topic, bs).WithShardingKey(key).WithTag(tag)
	if delayLvl != 0 {
		pMsg.WithDelayTimeLevel(delayLvl)
	}
	// fixme consumer disorder producer?
	_, err = p.producer.SendSync(ctx, pMsg)
	return err
}

func (p *Producer) Close() error {
	return p.producer.Shutdown()
}

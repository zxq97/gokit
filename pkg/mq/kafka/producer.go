package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/mq"
)

var _ mq.Producer = (*Producer)(nil)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(conf *Config) (mq.Producer, error) {
	kfkConf := sarama.NewConfig()
	kfkConf.Producer.RequiredAcks = sarama.WaitForAll
	kfkConf.Producer.Retry.Max = 3
	kfkConf.Producer.Return.Successes = true
	kfkConf.Net.DialTimeout = defaultDialTimeout
	kfkConf.Net.ReadTimeout = defaultReadTimeout
	kfkConf.Net.WriteTimeout = defaultWriteTimeout
	producer, err := sarama.NewSyncProducer(conf.Addr, kfkConf)
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: producer,
	}, nil
}

func (p *Producer) SendMessage(ctx context.Context, topic, key, tag string, msg proto.Message, _ int) error {
	bs, err := mq.WarpMessage(ctx, tag, msg)
	_, _, err = p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(bs),
	})
	return err
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

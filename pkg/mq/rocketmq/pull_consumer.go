package rocketmq

import "github.com/zxq97/gokit/pkg/mq"

var _ mq.Consumer = (*PullConsumer)(nil)

type PullConsumer struct {
}

func (c *PullConsumer) Start() error {
	return nil
}

func (c *PullConsumer) Stop() error {
	return nil
}

package mq

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
)

type Consumer interface {
	Start() error
	Stop() error
}

type HandleFunc func(ctx context.Context, msg *MqMessage) error

type Option func(*ConsumerMeta)

func WithName(name string) Option {
	return func(c *ConsumerMeta) {
		c.Name = name
	}
}

func WithNPoll(n int) Option {
	return func(c *ConsumerMeta) {
		c.NPoll = n
	}
}

func WithNProc(n int) Option {
	return func(c *ConsumerMeta) {
		c.NProc = n
	}
}

func WithProcTimeout(d time.Duration) Option {
	return func(c *ConsumerMeta) {
		c.ProcTimeout = d
	}
}

type ConsumerMeta struct {
	Name        string
	Fn          HandleFunc
	NPoll       int
	NProc       int
	ProcTimeout time.Duration
	Done        chan struct{}
}

func Unmarshal(message []byte) (*MqMessage, error) {
	msg := &MqMessage{}
	err := proto.Unmarshal(message, msg)
	return msg, err
}

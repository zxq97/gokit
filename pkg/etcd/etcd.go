package etcd

import (
	"time"

	"go.etcd.io/etcd/client/v3"
)

type Config struct {
	Addr []string `yaml:"addr"`
	TTL  int      `yaml:"ttl"`
}

func WithEndpoints(addr []string) Option {
	return func(c *clientv3.Config) {
		c.Endpoints = addr
	}
}

func WithDialTimeout(d time.Duration) Option {
	return func(c *clientv3.Config) {
		c.DialTimeout = d
	}
}

type Option func(*clientv3.Config)

func NewEtcd(conf *Config) (*clientv3.Client, error) {
	opts := []Option{WithEndpoints(conf.Addr)}
	if conf.TTL != 0 {
		opts = append(opts, WithDialTimeout(time.Duration(conf.TTL)*time.Second))
	}
	return NewEtcdByOption(opts...)
}

func NewEtcdByOption(opts ...Option) (*clientv3.Client, error) {
	c := clientv3.Config{}
	for _, o := range opts {
		o(&c)
	}
	return clientv3.New(c)
}

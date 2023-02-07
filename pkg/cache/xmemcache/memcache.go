package xmemcache

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Config struct {
	Addr        []string `yaml:"addr"`
	MaxIdleConn int      `yaml:"max_idle_conn"`
	Timeout     int      `yaml:"timeout"`
}

func WithTimeout(d time.Duration) Option {
	return func(c *memcache.Client) {
		c.Timeout = d
	}
}

func WithMaxIdleConn(cnt int) Option {
	return func(c *memcache.Client) {
		c.MaxIdleConns = cnt
	}
}

type Option func(*memcache.Client)

func NewXMemcache(conf *Config) *memcache.Client {
	opt := []Option{}
	if conf.MaxIdleConn != 0 {
		opt = append(opt, WithMaxIdleConn(conf.MaxIdleConn))
	}
	if conf.Timeout != 0 {
		opt = append(opt, WithTimeout(time.Duration(conf.Timeout)*time.Second))
	}
	return NewMemcache(conf.Addr, opt...)
}

func NewMemcache(addr []string, opts ...Option) *memcache.Client {
	c := memcache.New(addr...)
	for _, o := range opts {
		o(c)
	}
	return c
}

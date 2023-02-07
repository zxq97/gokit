package xredis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisTypeCluster = "cluster"
)

type Config struct {
	Addr         []string `yaml:"addr"`
	DB           int      `yaml:"db"`
	Type         string   `yaml:"type"`
	DialTimeout  int      `yaml:"dial_timeout"`
	ReadTimeout  int      `yaml:"read_timeout"`
	WriteTimeout int      `yaml:"write_timeout"`
}

type XRedis struct {
	redis.Cmdable
}

func WithClusterAddr(addr []string) ClusterOption {
	return func(o *redis.ClusterOptions) {
		o.Addrs = addr
	}
}

func WithClusterDialTimeout(d time.Duration) ClusterOption {
	return func(o *redis.ClusterOptions) {
		o.DialTimeout = d
	}
}

func WithClusterReadTimeout(d time.Duration) ClusterOption {
	return func(o *redis.ClusterOptions) {
		o.ReadTimeout = d
	}
}

func WithClusterWriteTimeout(d time.Duration) ClusterOption {
	return func(o *redis.ClusterOptions) {
		o.WriteTimeout = d
	}
}

func WithAddr(addr string) Option {
	return func(o *redis.Options) {
		o.Addr = addr
	}
}

func WithDB(db int) Option {
	return func(o *redis.Options) {
		o.DB = db
	}
}

func WithDialTimeout(d time.Duration) Option {
	return func(o *redis.Options) {
		o.DialTimeout = d
	}
}

func WithReadTimeout(d time.Duration) Option {
	return func(o *redis.Options) {
		o.ReadTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) Option {
	return func(o *redis.Options) {
		o.WriteTimeout = d
	}
}

type ClusterOption func(*redis.ClusterOptions)

type Option func(*redis.Options)

func NewXRedis(conf *Config) *XRedis {
	switch conf.Type {
	case redisTypeCluster:
		opts := []ClusterOption{WithClusterAddr(conf.Addr)}
		if conf.DialTimeout != 0 {
			opts = append(opts, WithClusterDialTimeout(time.Duration(conf.DialTimeout)*time.Second))
		}
		if conf.ReadTimeout != 0 {
			opts = append(opts, WithClusterReadTimeout(time.Duration(conf.ReadTimeout)*time.Second))
		}
		if conf.WriteTimeout != 0 {
			opts = append(opts, WithClusterWriteTimeout(time.Duration(conf.WriteTimeout)*time.Second))
		}
		return &XRedis{NewRedisClusterClient(opts...)}
	default:
		opts := []Option{WithAddr(conf.Addr[0])}
		if conf.DB != 0 {
			opts = append(opts, WithDB(conf.DB))
		}
		if conf.DialTimeout != 0 {
			opts = append(opts, WithDialTimeout(time.Duration(conf.DialTimeout)*time.Second))
		}
		if conf.ReadTimeout != 0 {
			opts = append(opts, WithReadTimeout(time.Duration(conf.ReadTimeout)*time.Second))
		}
		if conf.WriteTimeout != 0 {
			opts = append(opts, WithWriteTimeout(time.Duration(conf.WriteTimeout)*time.Second))
		}
		return &XRedis{NewRedisClient(opts...)}
	}
}

func NewRedisClient(opts ...Option) *redis.Client {
	opt := &redis.Options{}
	for _, o := range opts {
		o(opt)
	}
	return redis.NewClient(opt)
}

func NewRedisClusterClient(opts ...ClusterOption) *redis.ClusterClient {
	opt := &redis.ClusterOptions{}
	for _, o := range opts {
		o(opt)
	}
	return redis.NewClusterClient(opt)
}

package xredis

import "github.com/go-redis/redis/v8"

const (
	redisTypeCluster = "cluster"
)

type Config struct {
	Addr []string `yaml:"addr"`
	DB   int      `yaml:"db"`
	Type string   `yaml:"type"`
}

type XRedis struct {
	redis.Cmdable
}

func NewRedis(conf *Config) *XRedis {
	switch conf.Type {
	case redisTypeCluster:
		return &XRedis{
			redis.NewClusterClient(&redis.ClusterOptions{
				Addrs: conf.Addr,
			}),
		}
	default:
		return &XRedis{
			redis.NewClient(&redis.Options{
				Addr: conf.Addr[0],
				DB:   conf.DB,
			}),
		}
	}
}

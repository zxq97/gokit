package etcd

import (
	"time"

	"go.etcd.io/etcd/client/v3"
)

type Config struct {
	Addr []string `yaml:"addr"`
	TTL  int      `yaml:"ttl"`
}

func NewEtcd(conf *Config) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   conf.Addr,
		DialTimeout: time.Duration(conf.TTL) * time.Second,
	})
}

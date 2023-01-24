package xmemcache

import "github.com/bradfitz/gomemcache/memcache"

type Config struct {
	Addr []string `yaml:"addr"`
}

func NewMemcache(conf *Config) *memcache.Client {
	return memcache.New(conf.Addr...)
}

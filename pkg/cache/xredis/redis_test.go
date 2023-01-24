package xredis

import (
	"context"
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func getRedisClient() (*XRedis, error) {
	conf := &Config{}
	err := config.LoadYaml("../../../internal/yaml/redis.yaml", conf)
	if err != nil {
		return nil, err
	}
	log.Println(conf.Addr, conf.DB)
	return NewRedis(conf), nil
}

func TestNewRedis(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	val, err := client.HGet(context.TODO(), "k", "h").Result()
	if err != nil {
		t.Error(err)
	}
	log.Println(val)
}

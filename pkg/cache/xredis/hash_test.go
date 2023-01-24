package xredis

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestXRedis_HGetXEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	res, err := client.HGetXEX(context.TODO(), "h", "k", time.Minute)
	log.Println(res, err)
}

func TestXRedis_HMIncrByEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.HMIncrByEX(context.TODO(), "h", map[string]int64{"h": 1}, time.Minute)
	log.Println(err)
}

func TestXRedis_HIncrByXEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.HIncrByXEX(context.TODO(), "h", "k", 1, time.Minute)
	log.Println(err)
}

func TestXRedis_HMGetXEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	res, err := client.HMGetXEX(context.TODO(), "h", time.Minute, "k")
	log.Println(res, err)
}

func TestXRedis_HMSetEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.HMSetEX(context.TODO(), "h", map[string]interface{}{"k": "2"}, time.Minute)
	log.Println(err)
}

func TestXRedis_HMSetXEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.HMSetXEX(context.TODO(), "h", map[string]interface{}{"k": 1}, time.Minute)
	log.Println(err)
}

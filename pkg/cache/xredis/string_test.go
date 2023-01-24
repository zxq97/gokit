package xredis

import (
	"context"
	"testing"
	"time"
)

func TestXRedis_IncrByXEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.IncrByXEX(context.TODO(), "k", 1, time.Minute)
	if err != nil {
		t.Error(err)
	}
}

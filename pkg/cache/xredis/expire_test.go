package xredis

import (
	"context"
	"testing"
	"time"
)

func TestXRedis_ExistEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.ExistEX(context.TODO(), "k", time.Second)
	if err != nil {
		t.Error(err)
	}
}

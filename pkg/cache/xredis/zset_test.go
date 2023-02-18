package xredis

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestXRedis_ZAddEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.ZAddEX(context.TODO(), "z", []*redis.Z{{Member: "a", Score: 1}}, time.Minute)
	if err != nil {
		t.Error(err)
	}
}

func TestXRedis_ZAddXEX(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	err = client.ZAddXEX(context.TODO(), "z", []*redis.Z{{Member: "a", Score: 1}}, time.Minute)
	if err != nil {
		t.Error(err)
	}
}

func TestXRedis_ZRevPageByMember(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.ZRevPageByMember(context.TODO(), "z", "v", 2)
	log.Println(res, err)
}

func TestXRedis_ZRevPageByMemberWithScores(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.ZRevPageByMemberWithScores(context.TODO(), "z", "g", 2)
	log.Println(res, err)
}

func TestXRedis_ZRevRangeByMember(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	res, err := client.ZRevRangeByMember(context.TODO(), "z", "f", 1)
	log.Println(res, err)
}

func TestXRedis_ZRevRangeByMemberWithScores(t *testing.T) {
	client, err := getRedisClient()
	if err != nil {
		t.Error(err)
	}
	res, err := client.ZRevRangeByMemberWithScores(context.TODO(), "z", "d", 1)
	log.Println(res, err)
}

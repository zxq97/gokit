package xredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (xr *XRedis) ExistEX(ctx context.Context, key string, ttl time.Duration) error {
	ok, err := xr.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	} else if !ok {
		return redis.Nil
	}
	return nil
}

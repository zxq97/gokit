package xredis

import (
	"context"
	"time"
)

func (xr *XRedis) IncrByXEX(ctx context.Context, key string, incr int64, ttl time.Duration) error {
	if err := xr.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return xr.IncrBy(ctx, key, incr).Err()
}

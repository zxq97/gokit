package xredis

import (
	"context"
	"time"
)

func (xr *XRedis) HIncrByXEX(ctx context.Context, key, field string, incr int64, ttl time.Duration) error {
	if err := xr.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return xr.HIncrBy(ctx, key, field, incr).Err()
}

func (xr *XRedis) HMIncrByEX(ctx context.Context, key string, fieldMap map[string]int64, ttl time.Duration) error {
	pipe := xr.Pipeline()
	for k, v := range fieldMap {
		pipe.HIncrBy(ctx, key, k, v)
	}
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

func (xr *XRedis) HMSetEX(ctx context.Context, key string, fieldMap map[string]interface{}, ttl time.Duration) error {
	pipe := xr.Pipeline()
	pipe.HMSet(ctx, key, fieldMap)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

func (xr *XRedis) HMSetXEX(ctx context.Context, key string, fieldMap map[string]interface{}, ttl time.Duration) error {
	if err := xr.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return xr.HMSet(ctx, key, fieldMap).Err()
}

func (xr *XRedis) HMGetXEX(ctx context.Context, key string, ttl time.Duration, field ...string) ([]interface{}, error) {
	if err := xr.ExistEX(ctx, key, ttl); err != nil {
		return nil, err
	}
	return xr.HMGet(ctx, key, field...).Result()
}

func (xr *XRedis) HGetXEX(ctx context.Context, key, field string, ttl time.Duration) (string, error) {
	if err := xr.ExistEX(ctx, key, ttl); err != nil {
		return "", err
	}
	return xr.HGet(ctx, key, field).Result()
}

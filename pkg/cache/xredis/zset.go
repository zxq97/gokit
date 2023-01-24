package xredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zxq97/gokit/pkg/cast"
)

var (
	zRevRangeByMemberScript = redis.NewScript(`
        local rank = redis.call("ZRevRank", KEYS[1], ARGV[1])
        if not rank then
            rank = 0
		else
			rank = rank + 1
        end
        return redis.call("ZRevRange", KEYS[1], rank, rank + ARGV[2])
    `)

	zRevRangeByMemberWithScoresScript = redis.NewScript(`
		local rank = redis.call("ZRevRank", KEYS[1], ARGV[1])
		if not rank then
			rank = 0
		else
			rank = rank + 1
		end
		return redis.call("ZRevRange", KEYS[1], rank, rank + ARGV[2], 'withscores')
	`)
)

func (xr *XRedis) ZAddXEX(ctx context.Context, key string, zs []*redis.Z, ttl time.Duration) error {
	if err := xr.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return xr.ZAddEX(ctx, key, zs, ttl)
}

func (xr *XRedis) ZAddEX(ctx context.Context, key string, zs []*redis.Z, ttl time.Duration) error {
	pipe := xr.Pipeline()
	pipe.ZAdd(ctx, key, zs...)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

func (xr *XRedis) ZRevRangeByMember(ctx context.Context, key string, member interface{}, offset int64) ([]int64, error) {
	res, err := zRevRangeByMemberScript.Run(ctx, xr, []string{key}, member, offset).Result()
	if err != nil {
		return nil, err
	}
	val, ok := res.([]interface{})
	if !ok || len(val) == 0 {
		return nil, redis.Nil
	}
	ids := make([]int64, 0, offset)
	for _, v := range val {
		id := v.(string)
		ids = append(ids, cast.ParseInt(id, 0))
	}
	return ids, nil
}

func (xr *XRedis) ZRevRangeByMemberWithScores(ctx context.Context, key string, member interface{}, offset int64) ([]*redis.Z, error) {
	res, err := zRevRangeByMemberWithScoresScript.Run(ctx, xr, []string{key}, member, offset).Result()
	if err != nil {
		return nil, err
	}
	val, ok := res.([]interface{})
	if !ok || len(val) == 0 || len(val)&1 != 0 {
		return nil, redis.Nil
	}
	zs := make([]*redis.Z, 0, len(val)>>1)
	for i := 0; i < len(val); i += 2 {
		id := val[i+1].(string)
		zs = append(zs, &redis.Z{Member: val[i], Score: float64(cast.ParseInt(id, 0))})
	}
	return zs, nil
}

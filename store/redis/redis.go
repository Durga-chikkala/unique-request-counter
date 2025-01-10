package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) Redis {
	return Redis{rdb: rdb}
}

func (r *Redis) RequestCount(ctx context.Context) int64 {
	cmd := r.rdb.Get(ctx, "unique_request")

	result, err := cmd.Int64()
	if err != nil {
		return 0
	}

	return result
}

func (r *Redis) Flush(ctx context.Context) bool {
	cmd := r.rdb.FlushDB(ctx)

	err := cmd.Err()
	if err != nil {
		return false
	}

	return true
}

func (r *Redis) LockId(ctx context.Context, id int64) bool {
	ok, _ := r.rdb.SetNX(ctx, fmt.Sprintf("lock:%d", id), "locked", 1*time.Minute).Result()
	if ok {
		return true
	}

	return false
}

func (r *Redis) IncrementCount(ctx context.Context) {
	r.rdb.Incr(ctx, "unique_request")
}

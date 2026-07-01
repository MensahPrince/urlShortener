package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Allow(ctx context.Context, ip string, limit int) (bool, error) {
	count, err := rdb.Incr(ctx, ip).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		rdb.Expire(ctx, ip, time.Minute)
	}

	return count <= int64(limit), nil
}

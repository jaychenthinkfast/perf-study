package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})
	_ = rdb.FlushDB(ctx).Err()
	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(10))
	if err != nil {
		panic(err)
	}

	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)

}

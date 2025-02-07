package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/go-redis/redis/v8" // 假设使用 go-redis
)

var (
	rdb          = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	requestGroup singleflight.Group
	ctx          = context.Background()
)

func getValueFromRedis(key string) (string, error) {
	// 查询 Redis 缓存
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("Cache miss, querying DB...")
		// 使用 singleflight.Group 防止并发缓存重建
		v, err, shared := requestGroup.Do(key, func() (interface{}, error) {
			fmt.Println("DB query started")
			// 模拟从数据库查询
			time.Sleep(2 * time.Second)
			result := "data_from_db"
			// 存入 Redis
			err := rdb.Set(ctx, key, result, time.Minute).Err()
			return result, err
		})

		if err != nil {
			return "", err
		}
		if shared {
			fmt.Println("Result was shared among requests")
		}
		return v.(string), nil
	} else if err != nil {
		return "", err
	}
	fmt.Println("Cache hit")
	return value, nil
}

func main() {
	key := "test_key"
	var wg sync.WaitGroup
	// 模拟高并发请求
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			value, err := getValueFromRedis(key)
			if err != nil {
				log.Printf("Goroutine %d: error: %v\n", id, err)
				return
			}
			log.Printf("Goroutine %d: got value: %s\n", id, value)
		}(i)
	}
	wg.Wait()
}

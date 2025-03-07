package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// bloomFilter 是一个使用 RedisBloom 布隆过滤器的函数
func bloomFilter(ctx context.Context, rdb *redis.Client) {
	// 添加元素到布隆过滤器
	inserted, err := rdb.Do(ctx, "BF.ADD", "bf_key", "item0").Bool()
	if err != nil {
		panic(err)
	}
	if inserted {
		fmt.Println("item0 was inserted")
	}

	// 检查元素是否存在
	for _, item := range []string{"item0", "item1"} {
		exists, err := rdb.Do(ctx, "BF.EXISTS", "bf_key", item).Bool()
		if err != nil {
			panic(err)
		}
		if exists {
			fmt.Printf("%s does exist\n", item)
		} else {
			fmt.Printf("%s does not exist\n", item)
		}
	}
}

// 初始化 Redis 客户端
func newRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:63790", // Redis 地址
		Password: "",                // Redis 密码，默认无密码
		DB:       0,                 // 使用默认 DB
	})
	return rdb
}

// 主函数，用于手动运行
func main() {
	ctx := context.Background()
	rdb := newRedisClient()

	// 确保 Redis 连接正常
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	defer rdb.Close()

	// 清空之前的 bf_key（可选）
	rdb.Del(ctx, "bf_key")

	// 调用布隆过滤器函数
	bloomFilter(ctx, rdb)
}

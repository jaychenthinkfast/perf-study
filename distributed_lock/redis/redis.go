package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

func main() {
	// 创建 Redis 客户端
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client)

	// 初始化 Redsync
	rs := redsync.New(pool)

	// 创建分布式锁
	mutex := rs.NewMutex("my-distributed-lock",
		redsync.WithExpiry(10*time.Second),           // 锁过期时间 10 秒
		redsync.WithTries(5),                         // 最多重试 5 次
		redsync.WithRetryDelay(500*time.Millisecond), // 每次重试间隔 500ms
	)

	// 上下文用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 获取锁
	fmt.Println("尝试获取锁...")
	if err := mutex.LockContext(ctx); err != nil {
		fmt.Printf("获取锁失败: %v\n", err)
		return
	}
	fmt.Println("成功获取锁")

	// 模拟业务逻辑
	go func() {
		// 手动续期，每 5 秒检查并延长锁
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if ok, err := mutex.ExtendContext(ctx); ok {
					fmt.Println("锁已续期")
				} else {
					fmt.Printf("续期失败: %v\n", err)
					return
				}
			}
		}
	}()

	// 模拟长时间任务
	time.Sleep(20 * time.Second)

	// 释放锁
	if ok, err := mutex.UnlockContext(ctx); ok {
		fmt.Println("成功释放锁")
	} else {
		fmt.Printf("释放锁失败: %v\n", err)
	}
}

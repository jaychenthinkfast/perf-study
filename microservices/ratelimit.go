package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// 创建每秒生成1个令牌，桶容量为2的限流器
	limiter := rate.NewLimiter(1, 2)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	for i := 0; i < 20; i++ {
		// 等待直到获取一个令牌
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Request", i, time.Now())
	}
}

package main

import (
	"context"
	"testing"
	"time"
)

// 测试函数
func TestBloomFilter(t *testing.T) {
	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 初始化 Redis 客户端
	rdb := newRedisClient()
	defer rdb.Close()

	// 检查 Redis 连接
	if err := rdb.Ping(ctx).Err(); err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}

	// 清空之前的 bf_key
	rdb.Del(ctx, "bf_key")

	// 调用布隆过滤器函数
	bloomFilter(ctx, rdb)

	// 验证结果
	existsItem0, err := rdb.Do(ctx, "BF.EXISTS", "bf_key", "item0").Bool()
	if err != nil {
		t.Fatalf("BF.EXISTS failed for item0: %v", err)
	}
	if !existsItem0 {
		t.Errorf("Expected item0 to exist, but it does not")
	}

	existsItem1, err := rdb.Do(ctx, "BF.EXISTS", "bf_key", "item1").Bool()
	if err != nil {
		t.Fatalf("BF.EXISTS failed for item1: %v", err)
	}
	if existsItem1 {
		t.Errorf("Expected item1 to not exist, but it does")
	}
}

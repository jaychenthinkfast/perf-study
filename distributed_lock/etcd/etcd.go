package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建会话（Session），绑定租约
	s, err := concurrency.NewSession(cli, concurrency.WithTTL(10)) // TTL 10秒
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// 创建分布式锁
	lock := concurrency.NewMutex(s, "/my-lock/")
	ctx := context.Background()

	// 获取锁
	if err := lock.Lock(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功获取锁！")

	// 模拟业务处理
	time.Sleep(5 * time.Second)

	// 释放锁
	if err := lock.Unlock(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("锁已释放！")
}

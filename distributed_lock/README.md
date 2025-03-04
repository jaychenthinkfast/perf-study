# 使用 etcd 实现分布式锁
## 背景
在分布式系统中，为了确保多个进程或节点不会同时访问共享资源，我们需要一种机制来实现互斥访问，这就是分布式锁的作用。
而etcd 作为高可用、强一致性的分布式键值存储系统，广泛用于配置管理、服务发现和分布式协调。其基于 Raft 共识算法，能够在集群中保证数据一致性，非常适合实现分布式锁。

## 实现原理
1. **Lease 机制（租约）**
   etcd 提供了租约（Lease）功能，客户端可以申请一个带有 TTL（生存时间）的租约。如果客户端在租约到期前未续约，锁会自动释放。这非常适合处理客户端宕机的情况，避免锁被永久持有。
2. **事务操作（Transaction）**
   etcd 的事务（txn）支持原子操作，可以通过比较键的版本号（revision）来判断锁是否已被占用，从而实现互斥性。
3. **Key 的版本号（Revision）**
   每当 etcd 中的键被创建或修改时，会生成一个全局唯一的 revision 号。通过比较 revision，可以判断锁的持有顺序，确保公平性。
4. **Watch 机制**
   如果锁被占用，客户端可以通过 Watch 监听锁的释放事件，一旦锁可用即可尝试获取。

## 基本流程
1. **获取锁**
   客户端尝试通过事务操作创建一个带有租约的键（如 /lock/my-lock）。
   如果键不存在（即 revision 为 0），则创建成功，客户端获得锁。
   如果键已存在，说明锁被其他客户端持有，当前客户端进入等待状态（可通过 Watch 监听）。
2. **持有锁**
   客户端通过定期续约（keep-alive）保持租约有效，确保锁不会因超时被释放。
3. **释放锁**
   客户端完成任务后，删除键或让租约自然过期，锁被释放。

这些设计保证了
* 锁的互斥性（只有一个客户端持有锁）
* 安全性（宕机时锁会自动释放）
* 公平性（按请求顺序获取锁）

## etcd安装
以 mac 为例
```
brew install etcd
brew services start etcd
```

## 代码示例

```
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
```
如果是 etcd集群环境Endpoints 需要填入所有集群节点信息

同时启动多个进程可观察到进程交替获取到锁进行业务处理。

## 实用建议
1. 设置合理的 TTL
   TTL 太短会导致频繁续约，增加负载；太长则可能延迟锁释放。建议根据业务需求设置为 5-30 秒。
2. 避免锁粒度过细
   锁的竞争越激烈，性能越差。尽量将锁范围控制在必要的最小粒度。
3. 处理网络分区
   在网络分区时，etcd 可能暂时不可用。建议在客户端实现重试机制（如指数退避），并设置超时。
4. 监控与报警
   使用 etcd 的内置 metrics（如 /metrics 端点），实用 prometheus等监控 QPS、延迟和租约状态，及时发现问题。
5. 测试竞争场景
   在开发阶段模拟高并发锁竞争，确保系统稳定性。


## 场景选择
如果需要高性能锁或大吞吐量，考虑 Redis；如果需要强一致性和复杂协调，etcd 是首选。

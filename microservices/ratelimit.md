# 限流
## 算法
| 常用限流算法 | 优势 | 不足 | 适用场景 |
|-------------|------|------|--------|
| **固定窗口算法（计数器算法）** | 简单易懂 | 请求流量分布不均时抗抖动能力弱，时间窗口边界限流不准 | 流量比较均匀且限流精度不高时 |
| **滑动窗口算法** | 相比固定窗口算法提升了限流精度 | 需要维护各个子窗口数据，实现复杂，面对流量波动，抗抖动能力弱 | 流量比较均匀且对限流精度有一定要求的场景，实践中很少见 |
| **漏桶算法** | 流量整形功能，即面对流量波动也能以固定速率处理请求 | 应对短时突发流量时，难以充分挖掘系统资源的潜力 | 流量绝对平滑的场景，比如对接第三方 API（电商支付中与多家银行系统对接） |
| **令牌桶算法** | 出色地应对短时突发流量 | 放弃了一定的流量整形功能，流量整体平滑但不是绝对平滑 | 流量整体平滑，允许一定短时突发流量的场景，比如微服务架构的上游服务限流，这是实践中得最多的限流算法 |

## 对象
本服务 - 端口（接口） - 实例（集群）
上游服务 - 端口（接口） - 实例（集群）

按需组合组装成 key 进行精准限流，区分实例主要是可能业务会按场景和保障等级按集群划分流量保障重点的业务场景

## 方式
| 限流方式   | 特点 | 优势 | 劣势 | 适用场景 |
|------------|------|------|------|----------|
| **集中式限流** | 依赖中心化组件（如 Redis），统一管理所有服务实例的限流 | - 能够精准控制整个系统的总流量<br>- 适用于多实例、分布式环境<br>- 便于动态调整限流策略 | - 需要额外的中心化组件，增加系统复杂度<br>- 依赖外部存储，存在网络延迟<br>- 高并发时，中心组件易成为瓶颈，影响稳定性 | - 适用于多实例、流量不均衡的分布式服务<br>- 需要全局流量管理的场景（如 API 网关、支付系统） |
| **单机本地限流** | 在每个服务实例内存中独立执行限流 | - 限流操作本地执行，延迟低、性能高<br>- 无需依赖外部组件，稳定性更好<br>- 适用于负载均衡较好的微服务架构 | - 无法保证全局流量的均衡控制<br>- 面对流量分布不均时，难以精准设定限流阈值 | - 适用于流量均衡的微服务架构<br>- 需要高效、低延迟的限流场景（如 Web 服务器、缓存服务） |

## 实现
### golang.org/x/time/rate 
实现了令牌桶算法
```go
package main

import (
    "context"
    "fmt"
    "time"

    "golang.org/x/time/rate"
)

func main() {
    // 创建每秒生成5个令牌，桶容量为10的限流器
    limiter := rate.NewLimiter(5, 10)

    for i := 0; i < 20; i++ {
        // 等待直到获取一个令牌
        err := limiter.Wait(context.Background())
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        fmt.Println("Request", i, time.Now())
    }
}

```
limiter.Wait 会阻塞直到获取到令牌。也可以设置带取消或者超时的上下文。
```go
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

```
也可以使用 limiter.Allow 方法非阻塞地检查是否有可用令牌。
```go
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/time/rate"
    "time"
)

func main() {
    // 创建一个每秒生成5个令牌，桶容量为10的限流器
    limiter := rate.NewLimiter(5, 10)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if limiter.Allow() {
            fmt.Fprintln(w, "Request allowed")
        } else {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
        }
    })

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

```
Allow 方法是非阻塞的，即使没有足够的令牌，它也会立即返回 false。因此，适用于需要立即知道请求是否被允许的场景，例如在高并发情况下快速拒绝超出限流的请求。

可以使用 SetLimit 和 SetBurst 方法动态调整限流器的速率和突发容量：

```go
limiter.SetLimit(10)    // 每秒10个令牌
limiter.SetBurst(20)    // 桶容量20

```

### go-zero 限流（并发控制）
https://go-zero.dev/docs/tutorials/service/governance/limiter

http压测 https://github.com/rakyll/hey

grpc压测 https://github.com/bojand/ghz



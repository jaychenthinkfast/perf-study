# perf
## 性能测试
[benchmark/README.md](benchmark/README.md)

# 大 key
大 key危害，解决：压缩、拆分

[bigkey/README.md](bigkey/README.md)

# 责任链
常用中间件

[chain_of_responsibility/README.md](chain_of_responsibility/README.md)

# 代码陷阱
循环迭代变量、json解析、协程泄露

[code_traps/README.md](code_traps/README.md)

# b+树
[db/b+.md](db/b+.md)

# 使用 etcd 实现分布式锁
[distributed_lock/README.md](distributed_lock/README.md)

# epoll
| **方面**         | **gnet**                            | **golang net**                    |
|------------------|-------------------------------------|-----------------------------------|
| **性能**         | 高，在高并发下更优（少 goroutine）   | 一般，高并发下资源占用高          |
| **资源占用**     | 低（事件驱动，少量 goroutine）       | 高（每个连接一个 goroutine）      |
| **设计哲学**     | 事件驱动，非阻塞，专注性能          | 阻塞 I/O，简单易用，通用性强      |
| **功能完整性**   | 核心功能，需自行扩展               | 全面支持（TLS、HTTP 等）          |
| **易用性**       | 需要理解事件模型，稍复杂            | 开箱即用，简单直接               |
| **适用场景**     | 高性能服务器（如 Redis、Haproxy）   | 通用网络应用（如 HTTP 服务器）    |
| **生态兼容性**   | 与标准库不兼容，迁移成本高          | 原生支持，生态丰富                |

关键差异：
* golang net 的简单性：基于 Go 的“goroutine 是轻量线程”的理念，每个连接一个 goroutine，适合快速开发和中小规模应用。但在连接数激增时（如数十万连接），内存和调度开销会显著增加。
* gnet 的高效性：通过事件循环和批量处理连接，适合高吞吐量、低延迟的场景。但它牺牲了部分通用性和易用性，开发者需要更多手动控制。
  (它不像 golang net 那样为每个连接分配一个 goroutine，而是用少量 goroutine 管理大量连接，减少了内存和 CPU 的开销。
  内置内存池和 goroutine 池，进一步优化资源使用。)

[epoll/README.md](epoll/README.md)


# errgroup
## 特征
是 waitgroup 的升级版，
* 支持返回错误，只返回第一个错误
* 支持并发控制，控制并发协程数
* 支持任务取消,在Go和 Wait 中触发，没有对外提供的接口

[errgroup/README.md](errgroup/README.md)

# 函数选项模式
函数选项模式是一种在 Go 语言中广泛使用的设计模式，它允许在创建结构体实例时，使用可选参数灵活地进行配置，而不会导致构造函数参数过长或过于复杂。

[functional_options_pattern/README.md](functional_options_pattern/README.md)

# 泛型

[generics/README.md](generics/README.md)

# 并发安全的 map(localcache)
* 并发原语
* bigcache
  * 分段 lock
  * []byte存储真实 data，减少 gc

[lock/README.md](lock/README.md)

# 日志
在哪记，记什么

[logs/README.md](logs/README.md)

# 微服务
* 熔断[microservice/breaker.md](microservices/breaker.md)
* 部署[microservice/deploy.md](microservices/deploy.md)
* 隔离[microservice/isolation.md](microservices/isolation.md)
* 限流[microservice/ratelimit.md](microservices/ratelimit.md)
* 重试[microservice/retry.md](microservices/retry.md)
* 超时[microservice/timeout.md](microservices/timeout.md)

# 性能测试
* cpu
  * http
  * runtime
* mem

[pprof](pprof)

# 项目布局
* 请求处理层。接受请求参数校验
* 业务逻辑层。可以调用通用层和 DAO 层
* 通用处理层。封装通用逻辑（内外部）复用
* 数据持久层。

[project_layout/README.md](project_layout/README.md)

# 反射
[reflection/README.md](reflection/README.md)

# singleflight
## 场景
singleflight 的最大优势在于避免了高并发时对相同资源的重复请求，提升了系统的性能和资源利用率。适用场景包括：

* 缓存系统： 避免缓存击穿，多个请求查询相同的缓存数据时，只会执行一次数据库查询或 API 调用。
* 数据库查询： 在高并发查询时，避免对数据库的重复查询。
* HTTP 请求： 避免对外部 API 的重复请求，减少网络带宽消耗。

[singleflight/README.md](singleflight/README.md)

## 性能优化
通过 pprof 火焰图确定瓶颈点，优化

1. map 值空结构体
2. map 预设容量
3. slic 预设容量
4. for range 改为 for index

[skill/README.md](skill/README.md)

通过
* strings.Builder(预分配)提升性能
* 避免使用fmt.Sprint改用 strconv.Itoa提升性能
* Str2Bytes可以使用unsafe.Pointer 达到 0 内存分配
* 实用高性能的序列化库 sonic提升性能，减少内存分配

[skill1/README.md]

## 协程池
[skill2/workerpool/README.md](skill2/workerpool/README.md)

## sync.pool
[skill2/pool/README.md](skill2/pool/README.md)

## 无锁栈
[skill2/lockfreestack/README.md](skill2/lockfreestack/README.md)

## syncmap
[syncmap/README.md](syncmap/README.md)

## 测试
* 单元测试
* 测试覆盖率
* mock

[test/README.md](test/README.md)

## trace分析
[trace/README.md](trace/README.md)

[trace1/http/README.md](trace1/http/README.md)

[trace1/runtime/READMD.md](trace1/runtime/READMD.md)
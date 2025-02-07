# singleflight
## singleflight.Group 的关键实现
* 内部映射 (m):

    * singleflight.Group 内部有一个映射（m），用来存储每个请求的状态和结果。
    * 每个请求用一个唯一的标识符（通常是请求的 key 或其他参数）作为键，存储请求的状态（正在执行或已完成）和结果。
* Do 方法：

    * 当多个请求请求相同的资源时，它们都会调用 Do 方法。Do 方法会先检查是否已有相同的请求在执行，如果有，后来的请求会等待执行完成；如果没有，Do 会启动一个新的请求。
    * Do 方法的第二个参数是一个函数，表示需要执行的实际工作（比如查询数据库或从 Redis 获取数据）。 
    * 该函数的返回值会通过 Do 方法传递给所有等待的请求。
## 代码细节
```go
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

type call struct {
	done chan struct{} // 标识该请求是否完成
	val  interface{}   // 结果值
	err  error         // 错误
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	// 如果有一个正在执行的请求，则等待它完成
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		// 等待已在执行的请求完成
		<-c.done
		return c.val, c.err
	}

	// 否则创建一个新的请求并记录
	c := &call{done: make(chan struct{})}
	g.m[key] = c
	g.mu.Unlock()

	// 执行实际工作
	c.val, c.err = fn()
	close(c.done)

	// 清理
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

```

## 场景
singleflight 的最大优势在于避免了高并发时对相同资源的重复请求，提升了系统的性能和资源利用率。适用场景包括：

* 缓存系统： 避免缓存击穿，多个请求查询相同的缓存数据时，只会执行一次数据库查询或 API 调用。
* 数据库查询： 在高并发查询时，避免对数据库的重复查询。
* HTTP 请求： 避免对外部 API 的重复请求，减少网络带宽消耗。
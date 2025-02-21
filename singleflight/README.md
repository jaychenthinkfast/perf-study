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
// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
    mu sync.Mutex       // protects m
    m  map[string]*call // lazily initialized
}

// call is an in-flight or completed singleflight.Do call
type call struct {
    wg sync.WaitGroup

  // These fields are written once before the WaitGroup is done
  // and are only read after the WaitGroup is done.
  val interface{}
  err error
  
  // These fields are read and written with the singleflight
  // mutex held before the WaitGroup is done, and are read but
  // not written after the WaitGroup is done.
  dups  int
  chans []chan<- Result
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
    g.mu.Lock()
    if g.m == nil {
        g.m = make(map[string]*call)
    }
    if c, ok := g.m[key]; ok {
        c.dups++
        g.mu.Unlock()
        c.wg.Wait()

        if e, ok := c.err.(*panicError); ok {
            panic(e)
        } else if c.err == errGoexit {
            runtime.Goexit()
        }
        return c.val, c.err, true
    }
    c := new(call)
    c.wg.Add(1)
	g.m[key] = c
    g.mu.Unlock()

    g.doCall(c, key, fn)
    return c.val, c.err, c.dups > 0
}

```

## 场景
singleflight 的最大优势在于避免了高并发时对相同资源的重复请求，提升了系统的性能和资源利用率。适用场景包括：

* 缓存系统： 避免缓存击穿，多个请求查询相同的缓存数据时，只会执行一次数据库查询或 API 调用。
* 数据库查询： 在高并发查询时，避免对数据库的重复查询。
* HTTP 请求： 避免对外部 API 的重复请求，减少网络带宽消耗。
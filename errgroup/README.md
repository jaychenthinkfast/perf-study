# errgroup
## 特征
是 waitgroup 的升级版，
* 支持返回错误，只返回第一个错误
* 支持并发控制，控制并发协程数
* 支持任务取消,在Go和 Wait 中触发，没有对外提供的接口

## 原理
### 结构
```
type token struct{}

type Group struct {
    cancel func(error) // 这个作用是為了前面说的 WithContext 而来的

    wg sync.WaitGroup // errGroup底层的阻塞等待功能，就是通过WaitGroup实现的

    sem chan token // 用于控制最大运行的协程数

    err     error // 最后在Wait方法中返回的error
    errOnce sync.Once // 用于安全的设置err
}
```

### WithContext
```
func WithContext(ctx context.Context) (*Group, context.Context) {
    ctx, cancel := withCancelCause(ctx)
    return &Group{cancel: cancel}, ctx // 生成有取消功能的context
}
```
### SetLimit
```
func (g *Group) SetLimit(n int) {
    g.sem = make(chan token, n) // 设置通道容量
}
```
### Go
```
func (g *Group) Go(f func() error) {
    if g.sem != nil {
        g.sem <- token{} // 通道满则阻塞，用来控制最大并发数
    }

    g.wg.Add(1)
    go func() {
        defer g.done() // 底层调用g.wg.Done()

        if err := f(); err != nil {
            g.errOnce.Do(func() { // 安全的设置err变量
                g.err = err
                if g.cancel != nil {
                    g.cancel(g.err) // 任务运行出错，调用g.cancel方法，用context控制其它任务中止运行
                }
            })
        }
    }()
}

func (g *Group) done() {
    if g.sem != nil {
        <-g.sem // 协程运行完，通道读
    }
    g.wg.Done()
}
```
### Wait
```
// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
    g.wg.Wait()
    if g.cancel != nil {
            g.cancel(g.err)
    }
    return g.err
}
```

## test
```
go test . -v  |  go test  -v   | go test -run=. -v           
=== RUN   TestWaitGroup
--- PASS: TestWaitGroup (5.00s)
=== RUN   TestErrHandle
Failured fetched all URLs.
--- PASS: TestErrHandle (5.00s)
=== RUN   TestCancel
Failured fetched all URLs.
--- PASS: TestCancel (5.00s)
=== RUN   TestLimitGNum
Failured fetched all URLs.
--- PASS: TestLimitGNum (5.92s)
PASS
ok      errgroup        21.398s
```

## 扩展
获取全部协程错误的实现
* 使用 chan  进行传输，在协程内 err->chan，在主协程内出 chan
* 使用err 切片➕锁
  * github.com/facebookgo/errgroup
  * github.com/vardius/gollback
* 使用 err 切片（固定 len)
  * getallerrs_test.go
# 反射

reflect.ValueOf(xx).Type() 等同  reflect.Typeof(xx) 均显示的不是最底层资源类型
如果需要获取底层资源类型，需reflect.ValueOf(xx).Kind()

## test
编译阶段无法提前察觉错误，只有在程序运行调用这个函数时才会报错
```go
go test -v
=== RUN   TestMax
--- FAIL: TestMax (0.00s)
panic: unsupported kind [recovered]
        panic: unsupported kind

goroutine 18 [running]:
testing.tRunner.func1.2({0xcdb0400, 0xc0001061f0})
        /usr/local/Cellar/go/1.23.4/libexec/src/testing/testing.go:1632 +0x230
testing.tRunner.func1()
        /usr/local/Cellar/go/1.23.4/libexec/src/testing/testing.go:1635 +0x35e
panic({0xcdb0400?, 0xc0001061f0?})
        /usr/local/Cellar/go/1.23.4/libexec/src/runtime/panic.go:785 +0x132
reflection.TestMax(0xc000124680?)
        /Users/chenjie/work/go/src/go123/perf-study/reflection/main_test.go:12 +0x6d
testing.tRunner(0xc000124680, 0xcdd2f30)
        /usr/local/Cellar/go/1.23.4/libexec/src/testing/testing.go:1690 +0xf4
created by testing.(*T).Run in goroutine 1
        /usr/local/Cellar/go/1.23.4/libexec/src/testing/testing.go:1743 +0x390
exit status 2
FAIL    reflection      0.691s

```

## benchmark
性能差距巨大
```go
go test -bench=. -benchmem -run=none
goos: darwin
goarch: amd64
pkg: reflection
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkRegular-8              1000000000               0.4108 ns/op          0 B/op          0 allocs/op
BenchmarkReflection-8           88916538                13.36 ns/op            0 B/op          0 allocs/op
PASS
ok      reflection      1.846s

```
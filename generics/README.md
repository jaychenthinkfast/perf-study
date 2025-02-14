# 泛型
## example
```go

// Max使用泛型来比较两个同类型的值（要求类型是可比较的），并返回较大的值
func Max[T int | float32](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	//类型隐式传入，由编译器自动推导
	fmt.Println(Max(1, 2))
	//类型显式传入
	fmt.Println(Max[float32](1.10, 2.2))
}
```

## benchmark
性能和非泛型性能相当，略优
```go
 go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: generics
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkGenerics-8     1000000000               0.3950 ns/op          0 B/op          0 allocs/op
BenchmarkRegular-8      1000000000               0.6536 ns/op          0 B/op          0 allocs/op
PASS
ok      generics        1.620s

```


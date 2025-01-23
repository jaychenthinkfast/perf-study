```
# -bench 表示需要benchmark运行的方法,.表示运行本目录所有Benchmark开头的方法
# -benchmem 显示与内存分配相关的详细信息
# -benchtime 设定每个基准测试用例的运行时间
# -cpuprofile 生成 CPU 性能分析文件
# -memprofile 生成内存性能分析文件
go test -bench='.' -benchmem -benchtime=10s -cpuprofile='cpu.prof' -memprofile='mem.prof'
```

输出

```
goos: darwin
goarch: amd64
pkg: benchmark
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
#从左到右分别表示benchmark函数、运行次数、单次运行消耗的时间、单次运行内存分配的字节数和次数
BenchmarkBytes2StrRaw-8         1000000000               6.266 ns/op           0 B/op          0 allocs/op
BenchmarkBytes2StrUnsafe-8      1000000000               0.3921 ns/op          0 B/op          0 allocs/op
PASS
ok      benchmark       7.951s

```

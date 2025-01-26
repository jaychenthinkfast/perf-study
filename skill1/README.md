benchmark
```
go test -run=none -bench=.  -benchmem  -benchtime=10s  -cpuprofile cpu4.prof
```
输出
```
goos: darwin
goarch: amd64
pkg: skill1
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkGenerateIdsRaw-8                   7670           1703403 ns/op         4196663 B/op       3744 allocs/op
BenchmarkGenerateIdsBuilder-8              74906            182518 ns/op           30081 B/op       2735 allocs/op
BenchmarkGenerateIdsStrconv-8             157072             76359 ns/op           23392 B/op       1901 allocs/op
BenchmarkGenerateIdsUnsafe-8              271431             43724 ns/op           11072 B/op        901 allocs/op
PASS
ok      skill1  54.246s

```
Raw->Builder->Strconv->Unsafe
* strings.Builder （strings.Join,strings.Replace 用到）
  * Strconv
    * Unsafe
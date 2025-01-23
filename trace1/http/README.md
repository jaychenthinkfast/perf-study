
不同于 pprof ，trace 可以记录协程运行的详细情况


```
curl "http://localhost:8888/debug/pprof/trace?seconds=30" > trace.out
```
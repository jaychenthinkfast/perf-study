# pprof cpu http
web : http://127.0.0.1:8888/debug/pprof/
![](image/pprof.png)
```
curl "http://127.0.0.1:8888/debug/pprof/profile?seconds=30" > profile.pprof
```
* profile
  * cpu
* allocs、heap
  * mem
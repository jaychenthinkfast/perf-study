```
go run main.go
go tool trace trace.out
```


* Goroutine analysis
  * 排查协程阻塞情况
  ![](image/goroutines.png)![](image/main.main.png)![](image/block.png)
* View trace by Proc
  * 处理器使用状况
  ![](image/proc.png)
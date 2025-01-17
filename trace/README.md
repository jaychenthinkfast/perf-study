# trace study
trace 工具可以记录 Go 程序运行时的所有事件，例如协程的创建、结束、阻塞、解除阻塞，系统调用的进入和退出，GC 开始、结束和 stopTheWorld 等等；
trace 会记录每个事件的纳秒级时间戳和调用栈信息。
后续通过 trace 可视化工具打开记录的事件信息，就可以追踪事件的前驱后置，以及分析事件相关的协程的调度、阻塞等情况。

场景：对于存在严重阻塞或者存在网络 IO 瓶颈的场景，这种因为等待而导致的性能不达标，CPU 占用和内存消耗可能都很小，pprof 分析无能为力时。
## v1
功能：生成曼德勃罗特图

相关命令行：
```
cd v2
make run //运行
make trace //可视化 trace
```

## v2
引入生产者消费者模型

## v3
通过 trace【View trace by proc】发现同时最多只有 2 个协程运行，同时【Goroutine analysis】调度等待时间较长，分析 chan 无缓存导致，增加缓存

## v4
通过 trace【View trace by proc】发现最初运行时有大量阻塞事件，同时【Synchronization blocking 】有阻塞出现，分析入队速度过慢导致，优化入队 

# sync.map
```
type Map struct {
    mu Mutex
    read atomic.Pointer[readOnly] // 无锁读map
    dirty map[any]*entry // 加锁读写map
    misses int
}

// readOnly is an immutable struct stored atomically in the Map.read field.
type readOnly struct {
    m       map[any]*entry
    amended bool // true if the dirty map contains some key not in m.
}
```
## 主要结构：
* read： 一个只读的映射，类型为 atomic.Pointer[readOnly]，用于无锁读取操作。
* dirty： 一个常规的映射，类型为 map[any]*entry，用于加锁的读写操作。
* mu： 一个互斥锁，类型为 sync.Mutex，用于保护对 dirty 映射的并发访问。
* misses： 一个整数，记录在 read 映射中未命中的次数。

## 工作原理：

### 读取操作：

首先尝试从 read 映射中获取值。
如果未命中，则增加 misses 计数，并在互斥锁 mu 的保护下，从 dirty 映射中查找。
### 写入操作：

在互斥锁 mu 的保护下进行。
如果 dirty 映射为 nil，则根据 read 映射创建一个新的 dirty 映射，并将 read 映射的 amended 字段设置为 true。
将新的键值对存储在 dirty 映射中。
### read 和 dirty 的转换：

当 misses 计数达到一定阈值时，dirty 映射会被提升为新的 read 映射，以包含最新的数据。
在某些情况下，read 映射也可能被降级为 dirty 映射，以处理并发写入。


## 设计优势：

### 读写分离： 
通过将读取操作与写入操作分离，sync.Map 在读取时避免了锁的开销，提高了性能。

### 延迟删除： 
删除操作并不会立即从映射中移除键值对，而是将其标记为已删除，实际的删除操作会在后续的维护过程中进行。

### 高效的并发访问：
通过结合无锁读取和加锁写入，sync.Map 在高并发场景下表现出色。

## 劣势
sync.map在写频繁或者读经常 miss 的情况下会发生遍历 read 到 dirty 的拷贝，
双 map结构也增加了内存占用，尤其在大数据量会有性能问题，
生产中还是分段锁 map适用居多，例如 [bigcache](../lock/README.md#扩展)

# map gc 
**key-value 需要满足非指针这个条件，key/value 的大小也不能超过 128 字节，
如果超过 128 字节，key-value 就会退化成指针，导致被 GC 扫描**

```
go test -v
```
```
=== RUN   TestSmallStruct
    main_test.go:64: size 5000000 GC duration: 1.829548ms
--- PASS: TestSmallStruct (3.86s)
=== RUN   TestBigStruct
    main_test.go:78: size 5000000 GC duration: 183.796178ms
--- PASS: TestBigStruct (2.06s)
```
# 并发安全的 map 
* sync.Mutex
  * 适用于写多｜读少
* sync.RWMutex
  * 适用于读多写少
* segment lock
  * 可降低锁粒度
  * 与每个值单独设置相比 gc对象少
* atomic.Value
  * 适用于数据量不大，定时全量更新
  ```
    func (v *Value) CompareAndSwap(old, new any) (swapped bool)
    func (v *Value) Load() (val any)
    func (v *Value) Store(val any)
    func (v *Value) Swap(new any) (old any)
  ```
  
## 扩展
[bigcache](https://github.com/allegro/bigcache)
* 分段锁，降低锁粒度
* 减少 gc
![bigcache-total-flow.png](image/bigcache-total-flow.png)
> https://pandaychen.github.io/2020/03/03/BIGCACHE-ANALYSIS/
> https://developer.aliyun.com/article/1444199
> https://blog.allegro.tech/2016/03/writing-fast-cache-service-in-go.html

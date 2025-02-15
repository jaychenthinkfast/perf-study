# range 
1.22 前
```go
    for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
```
实际上，循环迭代变量 v  的作用域覆盖整个循环，
每次迭代时，新值都会被赋给位于同一内存地址的变量  v。当协程中的闭包函数开始执行时，
如果循环已完成所有迭代，此时变量  v  留存的值便会是切片的最后一个元素  c。
正因如此，所有协程最终输出的结果都是字符 c。

## 解法 1
将  v  作为参数传递进匿名函数
```go
for _, v := range values {
    go func(u string) {
        fmt.Println(u)
        done <- true
    }(v)
}

```
## 解法 2
创建一个新的变量，其作用域仅在本次迭代，避免外部变量污染
```go
for _, v := range values {
    v := v // 创建一个新的 'v'，其作用域仅在本次迭代
    go func() {
        fmt.Println(v)
        done <- true
    }()
}
```

# json
```go
str := `{"name": "killianxu", "age": 18}`
	var mp map[string]interface{}

	json.Unmarshal([]byte(str), &mp)
	age := mp["age"].(int) // 报错：panic: interface conversion: interface {} is float64, not int
	fmt.Println(age)
```
int 序列化再反序列化需要 **float64去解析**

# channel
```go
func handle(timeout time.Duration) *Obj {
    ch := make(chan *Obj)
    go func() {
        result := fn() // 逻辑处理
        ch <- result   // block
    }()
    select {
    case result := <-ch:
        return result
    case <-time.After(timeout):
        return nil
    }
}
```
fn 长时间阻塞超过 timeout 会导致主协程 退出，导致 ch<-result 阻塞
可调整 channel 增加 buffer 比如 ch := make(chan *Obj, 1)

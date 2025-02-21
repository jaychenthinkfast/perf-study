# 重试设置
一般设置 2-3 次，重试总时长应小于上游服务调用本服务超时时间，避免无效重试

不应该生硬地编码在代码里，而应该采用远程配置的方式。这是因为这些数值需要根据线上实际运行状况灵活调整。

通过在 HTTP 协议头部，或是在 RPC 的 Request 中添加扩展字段的方式，实现对重试请求的有效标识。在执行重试操作时，携带上这个扩展字段，并在整个服务链路中进行透明传递。这样一来，链路上的各个服务就能根据这个特定的扩展字段，准确识别出来这是一个重试请求。
```go
// 请求处理函数，从HTTP头部获取重试字段，放入context，调用下一跳服务
func handler(w http.ResponseWriter, r *http.Request) {
    // 从请求头获取x - retry - flag标识
    retryFlag := r.Header.Get("x-retry-flag")
    // 将这个标识设置到context中，这样这个服务的后续处理逻辑都能从这个context里面获取重试标识
    ctx := context.WithValue(r.Context(), "x-retry-flag", retryFlag)

    // 构建下一跳服务的请求
    nextReq, err := http.NewRequestWithContext(ctx, "GET", "http://next-hop-service", nil)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // 从context中获取x - retry - flag标识，并设置到请求头
    if val, ok := ctx.Value("x-retry-flag").(string); ok {
        nextReq.Header.Set("x-retry-flag", val)
    }

    // 这里应该是实际发起请求到下一跳服务的逻辑，这里简单打印请求信息
    fmt.Printf("Sending request to next hop with retry flag: %s\n", retryFlag)
    // 实际可以使用http.Client发起请求，例如：
    // client := &http.Client{}
    // resp, err := client.Do(nextReq)
    //... 处理响应和错误
}
```
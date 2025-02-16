#log
## 输出位置
* 服务启动阶段
* 请求处理时，出入参
* 分支处理时
* 外部交互
* 定时后台任务
* 异常场景
## 记录内容
* 事件描述
* 事件状态
* 触发原因
* 关键上下文
* 函数调用堆栈信息

例子：
```go
2024 - 01 - 01 12:34:56.789 [192.168.1.100] [req - 123456] INFO - user login attempt success，user_id:123
2024 - 01 - 01 12:35:00.123 [192.168.1.100] [req - 123457] WARN - user login attempt failed due to invalid password，user_id:123，pwd:123
2024 - 01 - 01 12:35:10.567 [192.168.1.100] [req - 123458] ERROR - file upload failed due to insufficient disk space，file_name:example.txt, file_size:1024KB
Stack trace:
github.com/your - project/pkg/upload.(*Uploader).Upload
        /path/to/your - project/pkg/upload/uploader.go:256
github.com/your - project/cmd/app.main
        /path/to/your - project/cmd/app/main.go:46
```
## github.com/pkg/errors
一个用于增强 Go 语言错误处理的库，它提供了错误包装、堆栈跟踪和错误原因提取等功能。
以下是一个示例，展示如何使用 errors.Wrap 来包装错误，并使用 errors.Cause 获取原始错误。
```go
// 模拟一个可能返回错误的函数
func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// 使用 errors.Wrap 包装错误，添加上下文信息
		return errors.Wrap(err, "failed to open file")
	}
	return nil
}

func main() {
	err := readFile("nonexistent.txt")
	if err != nil {
		// 打印错误信息，包含堆栈跟踪
		fmt.Printf("Error: %+v\n", err)

		// 获取原始错误
		cause := errors.Cause(err)
		fmt.Printf("Root cause: %v\n", cause)
	}
}
```
```go
go run main.go
```
输出
```go
Error: open nonexistent.txt: no such file or directory
failed to open file
main.readFile
        /Users/chenjie/work/go/src/perf-study/logs/main.go:14
main.main
        /Users/chenjie/work/go/src/perf-study/logs/main.go:20
runtime.main
        /usr/local/Cellar/go/1.23.4/libexec/src/runtime/proc.go:272
runtime.goexit
        /usr/local/Cellar/go/1.23.4/libexec/src/runtime/asm_amd64.s:1700
Root cause: open nonexistent.txt: no such file or directory

```
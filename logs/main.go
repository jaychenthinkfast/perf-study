package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

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

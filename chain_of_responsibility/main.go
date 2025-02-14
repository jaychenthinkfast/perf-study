package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志记录中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 计算请求处理时间
		duration := time.Since(startTime)
		// 获取请求的状态码
		status := c.Writer.Status()
		// 获取请求的客户端IP
		clientIP := c.ClientIP()
		// 获取请求的方法
		method := c.Request.Method
		// 获取请求的路径
		path := c.Request.URL.Path

		// 打印日志
		fmt.Printf("[%s] %s %s %s %d %s\n",
			startTime.Format("2006-01-02 15:04:05"),
			clientIP,
			method,
			path,
			status,
			duration,
		)
	}
}

func main() {
	// 创建 Gin 引擎实例
	r := gin.Default()

	// 注册全局中间件
	r.Use(LoggerMiddleware())

	// 定义一个简单的 GET 路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 启动 HTTP 服务
	r.Run(":8080")
}

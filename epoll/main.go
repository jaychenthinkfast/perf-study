package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// 用于读取客户端发送数据的缓冲区
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取客户端数据出错:", err)
			break
		}
		// 输出客户端发送的数据
		fmt.Printf("从客户端接收到: %s\n", buffer[:n])
		// 向客户端发送响应信息
		_, err = conn.Write([]byte("已收到你的消息\n"))
		if err != nil {
			fmt.Println("向客户端发送响应出错:", err)
			break
		}
	}
}

func main() {
	// 监听的地址和端口
	listenAddr := ":8888"
	// 创建监听的TCP套接字
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("监听出错:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("服务器正在监听 %s\n", listenAddr)
	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端连接出错:", err)
			continue
		}
		// 启动一个协程来处理客户端连接
		go handleConnection(conn)
	}
}

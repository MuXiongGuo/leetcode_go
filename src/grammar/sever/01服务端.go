package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(listener)
	// 阻塞等待用户连接
	conn, err := listener.Accept() // err被刷新了 只要有一个是定义的就可以
	if err != nil {
		fmt.Println(err)
		return
	}
	// 接收用户请求
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
}

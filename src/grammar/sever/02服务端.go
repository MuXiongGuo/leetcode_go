package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		if string(buf[:n-1]) == "exit" {
			fmt.Println(addr, " exit")
		}
		// 转换成大写进行输出
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
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
	for {
		conn, err := listener.Accept() // err被刷新了 只要有一个是定义的就可以
		if err != nil {
			fmt.Println(err)
			return
		}
		// 处理用户请求 每来一个用户就新建一个协程
		go HandleConn(conn)
	}
}

package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	buf := make([]byte, 2048)
	for {
		// 休息一下 再处理 模拟 大处理过程
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second * 5)
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		// 正则表达式去除转义字符
		if string(buf[:n-1]) == "exit" {
			fmt.Println(addr, " exit")
			return
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

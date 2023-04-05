package main

import (
	"fmt"
	"net"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
	// 发送数据
	_, err = conn.Write([]byte("r u ok?"))
	if err != nil {
		return
	}
}

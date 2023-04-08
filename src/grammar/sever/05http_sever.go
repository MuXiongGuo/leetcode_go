package main

// 用于包测试
import (
	"fmt"
	"log"
	"net"
)

func main() {
	// 创建监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Println(err)
		}
	}(listener)

	// 阻塞等待用户进入
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 接收客户端数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if n == 0 {
		log.Println(err)
	}
	fmt.Printf("#%v#", string(buf[:n])) // 打印报文格式
}

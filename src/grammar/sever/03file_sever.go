package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	buf := make([]byte, 1024*4)
	// 接收文件内容
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println(err)
			}
			return
		}
		f.Write(buf[:n])
	}
}
func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// 阻塞等待连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	fileName := string(buf[:n])

	// 回复ok
	conn.Write([]byte("ok"))

	// 接收文件内容
	ReceiveFile(fileName, conn)
}

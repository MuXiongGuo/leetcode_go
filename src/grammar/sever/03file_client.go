package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(f)
	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println(" 文件发送完毕")
			} else {
				fmt.Println(err)
			}
			return
		}
		conn.Write(buf[:n]) // 给服务器发送内容
	}
}
func main() {
	// input file name
	fmt.Println("input file name: ")
	var path string
	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// get file information
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// connect sever
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)
	// send file name
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println(err)
		return
	}
	// receive
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(buf[:n]) == "ok" {
		// start to send content
		SendFile(path, conn)
	}
}

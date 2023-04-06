package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 接收服务器回复的数据  使用协程就是为了能一边发送一边接收这个数据, 因为这是两个for循环 所以要并发才可以
	go func() {
		// 从键盘输入内容, 给服务器发送内容
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) // 从键盘读取内容放str中
			if err != nil {
				fmt.Println(err)
				return
			}
			conn.Write(str[:n])
		}

	}()
	// 切片缓存
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) // 接收服务器的请求
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n])) // 打印接收到的内容
	}
}

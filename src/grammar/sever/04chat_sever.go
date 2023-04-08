package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 聊天室本质就是给在的client转发消息
// 用哈希表来存储用户信息(用切片好像也可以)
// 客户端的信息统一用结构体保存
// 专门设置一个协程来传递信息(利用管道)

type Client struct {
	C    chan string // 用于发送数据的管道 自己配一个管道在结构体里
	Name string      // 用户名
	Addr string      // 网络地址
}

var onlinemap = make(map[string]Client)
var message = make(chan string) // 换成main函数里局部的可以吗??? 一定要全局变量吗

func HandleConn(conn net.Conn) {
	defer conn.Close()
	// 处理用户连接
	cliAddr := conn.RemoteAddr().String()

	// 创建结构体成员
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlinemap[cliAddr] = cli

	// 新开一个协程专门给客户端发送消息
	go WriteToClient(cli, conn)

	// 只告诉用户自己 他在哪里
	cli.C <- MakeMsg(cli, "I am here\n")

	// 广播某个人在线
	message <- MakeMsg(cli, "login\n")

	// 对方是否主动退出
	isQuit := make(chan bool)
	// 对方是否主动发送消息(超时太久的可以踢掉他)
	hasData := make(chan bool)

	// 新开一个协程用于用户发言以及离线处理
	go func() {
		buf := make([]byte, 2048)
		for {
			n, _ := conn.Read(buf) // 没发送东西的时候会阻塞, 如果退出也会有返回值的
			if n == 0 {            // 对方断开或者出问题了
				isQuit <- true
				return
			}
			// 添加功能 查询在线用户
			msg := string(buf[:n-1]) // 去掉额外的换行符(有能力的可以用正则表达式去掉)
			if msg == "who" {
				// 给当前用户发送所有在线的成员
				for _, tmp := range onlinemap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				// 添加重命名功能 eg:rename|mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlinemap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				// 转发消息
				message <- MakeMsg(cli, string(buf[:n]))
			}
			hasData <- true // 代表用户发送过数据
		}
	}()

	// 防止主进程结束
	for {
		select {
		case <-isQuit:
			msg := MakeMsg(cli, "exit")
			delete(onlinemap, cliAddr)
			fmt.Println(msg)
			message <- msg
			return
		case <-hasData: // 不用处理
		case <-time.After(10 * time.Second): // 60s后超时
			msg := MakeMsg(cli, "time out leave \n")
			delete(onlinemap, cliAddr)
			fmt.Println(msg)
			message <- msg
			return
		}
	}
}
func MakeMsg(cli Client, msg string) (buf string) {
	// 代码复用
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

func WriteToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { // range管道的用法???  (一直读管道内容 永远不会结束)
		conn.Write([]byte(msg))
	}
}

func Manager() {
	// 新开一个协程,转发消息,有消息来了后给 map
	for {
		msg := <-message // 没收到消息会阻塞
		for _, cli := range onlinemap {
			cli.C <- msg // 遍历map给每个成员发送刚才的消息
		}
	}
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(listener)

	// 新开一个协程,转发消息,有消息来了后给map
	go Manager()

	// 主协程,循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleConn(conn)
	}
}

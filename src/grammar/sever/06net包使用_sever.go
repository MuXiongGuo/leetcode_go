package main

import (
	"fmt"
	"net/http"
)

func HandConn(w http.ResponseWriter, r *http.Request) {
	// w，给客户端回复数据  r，读取客户端发送的数据
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body", r.Body)
	w.Write([]byte("hello go")) // 给客户端回复

}

func main() {
	// 注册处理函数，用户连接自动调用指定的处理函数
	http.HandleFunc("/mike.html", HandConn)

	// 监听绑定
	http.ListenAndServe(":8000", nil)

}

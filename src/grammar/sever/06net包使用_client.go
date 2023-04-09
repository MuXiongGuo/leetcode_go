package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(resp *http.Response) {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp)

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	fmt.Println("Body = ", resp.Body) // 读取到的是一个流(是一个流)要接收，直接打印就是个地址

	// 读取body内容
	var content string
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取结束", err)
			break
		}
		content += string(buf[:n])
		fmt.Println(string(buf[:n]))
	}
	select {}
}

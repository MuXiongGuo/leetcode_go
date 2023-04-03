package main

import (
	"fmt"
	"os"
)

func main() {
	fin, err := os.Open("./Mt4.txt")
	if err != nil {
		fmt.Println(err)
	}
	// defer 中的操作也要进行错误处理  太严谨了!
	defer func(fin *os.File) {
		err := fin.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fin)
	buf := make([]byte, 10)
	for {
		n, _ := fin.Read(buf) // 文件流不断的输入到buf中直到全部过一遍, 一次输入的量是尽量把buf输入满
		if n == 0 {           // 文件结束时n为0
			break
		}
	}
	fmt.Println(string(buf))
}

package main

import "fmt"

// 终端命令查看反汇编后的结果
// go build 00test.go
// go tool objdump -S -s "main.main" test.exe
const y = 1000

func main() {
	fmt.Println(y)
}

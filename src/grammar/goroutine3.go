package main

import (
	"fmt"
	"runtime"
)

func main() {
	// n := runtime.GOMAXPROCS(1)
	n := runtime.GOMAXPROCS(2) // 两个cpu来跑两个goroutine, 就会竞争输出结果了,  换成1可以自行感受
	fmt.Println("n = ", n)     // println 自带换行符
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

package main // 声明 main 包，表明当前是一个可执行程序

import "fmt" // 导入内置 fmt 包

func main() { // main函数，是程序执行的入口
	var x byte = 55
	var y uint8 = x
	fmt.Println(x + y)
	//fmt.Println("Hello Golang!") // 在终端打印 Hello Golang!
}

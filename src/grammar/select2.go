package main

import "fmt"

func main() {
	// 随机性
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c1 <- 55
	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	}
	// 发送数据
	c3 := make(chan int, 1)
	c4 := make(chan int, 1)
	select {
	case c3 <- 55:
		fmt.Println("c3")
	case c4 <- 66:
		fmt.Println("c4")
	}
	// 死锁了 只能输入一个
	//fmt.Println(<-c3)
	//fmt.Println(<-c4)
}

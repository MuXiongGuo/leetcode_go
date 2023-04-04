package main

import "fmt"

// channel 的学习和使用  数据接收是阻塞的
func main() {
	c := make(chan int) // 阻塞的, 发送和接收除非都同时准备好 否则就会被阻塞

	go func() {
		defer fmt.Println("子协程结束")
		fmt.Println("子协程正在运行")
		c <- 666
	}()

	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main 协程结束")
}

package main

import (
	"fmt"
	"time"
)

// 有缓冲空间但是 收发速度不同时小测试  即收的过快,但是发送来不及
// 超过容量会怎么样呢?   会触发阻塞而已 仍然按照之前的逻辑, 不会那么频繁的阻塞而已
func main() {
	c := make(chan int, 30)
	go func() {
		for i := 0; i < 300; i++ {
			c <- i
		}
		close(c)
	}()
	time.Sleep(2 * time.Second)
	for el := range c {
		fmt.Println(el)
	}
}

package main

import (
	"fmt"
	"time"
)

// 无缓冲的
func main() {
	c := make(chan int, 0)
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("son goroutine over")

		for i := 0; i < 10; i++ {
			fmt.Println("子协程运行开始")
			c <- i
			fmt.Println("子协程结束")
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		fmt.Println("main开始")
		num := <-c
		fmt.Println(num)
		fmt.Println("main结束")
	}
	fmt.Println("main over")
}

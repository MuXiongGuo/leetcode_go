package main

import "fmt"

// close channel
func main() {
	c := make(chan int)
	c2 := make(chan int, 20)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		for i := 0; i < 10; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for {
		// 不加close的话 读不到结束信号的,会一直卡在for循环里
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	for data := range c2 {
		fmt.Println(data)
	}
}

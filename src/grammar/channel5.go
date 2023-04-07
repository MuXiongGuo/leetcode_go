package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	// 只写通道
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i
	}
}

func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
	for {

	}
}

func main() {
	c := make(chan int)
	go printer(c)
	time.Sleep(3 * time.Second)
	counter(c)

}

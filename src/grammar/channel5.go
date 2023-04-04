package main

import "fmt"

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
}

func main() {
	c := make(chan int, 10)
	go counter(c)
	printer(c)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("time out")
				o <- true
			}
		}
	}()
	c <- 666
	time.Sleep(1 * time.Second)
	fmt.Println(<-o)
}

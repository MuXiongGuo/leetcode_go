package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) // 延时1秒
	}
}

func main() {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) // 延时1秒
	}
	//fmt.Printf("end")
}

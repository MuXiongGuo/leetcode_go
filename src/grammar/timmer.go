package main

import (
	"fmt"
	"time"
)

func main() {
	// 接收部分 <-timer.C 可以放在其他goroutine 实现其他功能
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Println("t1 is:", t1)

	t2 := <-timer1.C
	fmt.Println("t2 is:", t2)
}

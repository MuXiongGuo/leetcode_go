package main

import "time"

// 拷贝实参指针 导致实参被分配到堆
func test(p *int) {
	go func() {
		println(p)
	}()
	time.Sleep(time.Second * 1)
}

func main() {
	x := 100
	p := &x
	test(p)
}

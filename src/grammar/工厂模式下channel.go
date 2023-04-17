package main

import "sync"

// 工厂方法讲goroutine和channel绑定在一起

type receiver struct {
	sync.WaitGroup // 匿名字段嵌入
	data           chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)
	go func() {
		defer r.Done()
		for d := range r.data {
			println(d)
		}
	}()

	return r
}

func main() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data) // 对应 不再向for range 发送信息
	r.Wait()      // 对应 defer r.Done()
}

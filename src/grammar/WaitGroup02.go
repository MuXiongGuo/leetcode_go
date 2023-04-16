package main

import (
	"sync"
)

// 多处WaitGroup
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()
		println("2 done")
	}()

	go func() {
		wg.Done()
		println("1 done")
	}()

	wg.Wait()
	//time.Sleep(time.Second * 1)
	println("main done")
}

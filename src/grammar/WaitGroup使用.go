package main

import (
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//wg.Add(1)  // wg.Add 最好是在gofun外部 不要在内部 容易提前结束
		go func(id int) {
			wg.Add(1) // wg.Add 内部也行 不如外部, 虽然它实现了原子操作
			defer wg.Done()
			println(strconv.Itoa(id))
		}(i)
	}

	println("wait...")
	wg.Wait()
	println("done")
}

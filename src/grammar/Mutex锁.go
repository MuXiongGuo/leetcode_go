package main

import (
	"sync"
	"time"
)

// 使用锁要注意 不要拷贝 否则会失效
// 使用 *data 或者 将*Mutex 放在data里面(但要注意初始化)

type data struct {
	sync.Mutex
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data
	go func() {
		defer wg.Done()
		d.test("A")
	}()

	go func() {
		defer wg.Done()
		d.test("B")
	}()
	wg.Wait()
}

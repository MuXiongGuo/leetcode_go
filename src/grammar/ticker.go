package main

import (
	"fmt"
	"time"
)

// 间隔一定时间相关到发送信息

func main() {
	ticker := time.NewTicker(time.Second * 1)
	i := 0
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println(i)

			if i == 5 {
				ticker.Stop() // 手动停止
			}
		}
	}()

	for {

	}
}

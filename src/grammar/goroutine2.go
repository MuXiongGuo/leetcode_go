package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 30; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 30; i++ {
		runtime.Gosched() // 暂时让出线程, 之后还会再回到之前的位置
		fmt.Println("hello")
	}
}

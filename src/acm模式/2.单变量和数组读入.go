package main

import "fmt"

func main() {
	var t int
	for {
		var sum int
		fmt.Scan(&t)
		if t == 0 {
			break
		} else {
			a := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scan(&a[i])
				sum += a[i]
			}
		}
		fmt.Println(sum)
	}
}

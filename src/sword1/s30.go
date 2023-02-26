package main

import "fmt"

type MinStack struct {
	data []int
}

func main() {
	d := MinStack{}
	c := make([]int, 0)
	c = append(c, 8)
	d.data = append(d.data, 55)
	fmt.Println(d)
	fmt.Println(c)
}

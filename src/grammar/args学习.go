package main

import (
	"fmt"
)

func print1(s string, args ...int) {
	fmt.Println(s)
	fmt.Println(args)
}
func print2(a interface{}) {
	fmt.Println(a)
}
func print3(args ...interface{}) {
	fmt.Println(args)
}
func main() {
	print1("hello ", 9, 3, 99)
	print2([]int{33, 55, 33}) // 临时变量即可
	print3(55, "33", []int{999, 0012})
}

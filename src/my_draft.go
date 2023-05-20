package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{9, 4}
	s1, s2 = s2, s1
	fmt.Println(s1)
	fmt.Println(s2)
}

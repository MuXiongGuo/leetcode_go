package main

import "fmt"

func main() {
	// string 底层还是 指针而已 复制就复制!!! 随便复制
	s1 := []string{"hello", "world"}
	s2 := make([]string, 2)
	s2[0] = s1[0]
	fmt.Println(&s2[0])
	fmt.Println(&s1[0])
}

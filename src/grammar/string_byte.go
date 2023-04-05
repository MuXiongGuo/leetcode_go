package main

import (
	"fmt"
	"strings"
)

func main() {
	// 发生多次copy
	s := make([]byte, 10)
	s[0] = 88
	s[1] = 74
	z := strings.ToLower(string(s))
	x := []byte(z)
	fmt.Println(x)
}

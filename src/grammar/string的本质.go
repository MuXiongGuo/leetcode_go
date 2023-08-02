package main

import "fmt"

func main() {
	// byte是uint8 别名  rune 是 int32别名
	// string本质是一个byte数组 里面元素自然是数字啦

	//string 的本质
	//double := make(map[int]int)
	double := map[int]int{}
	//s := "eeq"
	//c := 'e'
	//double[33] = (s[0])
	fmt.Println(double[33])
	s := []byte{90, 76, 98}
	fmt.Println(string(s))
}

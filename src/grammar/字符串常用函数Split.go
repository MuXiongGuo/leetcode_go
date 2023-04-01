package main

import (
	"fmt"
	"strings"
)

// 许多常用字符串函数
// Contains, Join, Index, Repeat, Replace, Split, Trim, Fields

func main() {
	s := "a man a plan a canal panama"
	fmt.Printf("%q\n", strings.Split(s, "a"))
	// strings.Replace()  replace的实现时 具体过程看不懂源码
	x := strings.Trim(" !!@!xx! Actung !!!", "! @") // cutset 只是表示要去除的集合, 顺序并不重要
	fmt.Println(x)
}

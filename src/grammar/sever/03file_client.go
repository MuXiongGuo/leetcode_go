package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("useage: xxx file")
		return
	}

	fileName := list[1]
	info, err := os.Stat(fileName) // 获取文件属性
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file name is ", info.Name())
	fmt.Println("file size is ", info.Size())
}

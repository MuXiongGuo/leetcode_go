package main

import "os"

type data struct {
	os.File
	int
	age int
}

func main() {
	// 匿名字段不需要带包名
	d := data{
		File: os.File{},
		int:  0,
		age:  0,
	}
	d.File = os.File{}
	d.int = 55
	d.age = 22
	//d.File.Name()
}

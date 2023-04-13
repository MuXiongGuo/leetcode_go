package main

import (
	"fmt"
)

func main() {
	type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{"test.txt", struct {
		owner int
		perm  int
	}{0, 0777}}
	// no!
	//f := file{
	//	name: "test.txt",
	//	attr: {
	//		owner: 0,
	//		perm:  0777,
	//	},
	//}
	fmt.Println(xx)
	fmt.Println(f)
}

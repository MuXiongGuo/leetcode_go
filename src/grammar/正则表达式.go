package main

import (
	"fmt"
	"regexp"
)

func main() {
	regexp1, _ := regexp.Compile("abcdefghijklmnopqrstuvwxyz")
	x := "dwoauo2331sa\n"
	fmt.Println(regexp1.FindString(x))
}

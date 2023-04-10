package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile("w.*w")
	x := "wdawjmygamedw\"ajio?p"
	z := re.FindString(x)
	fmt.Println(z)
}

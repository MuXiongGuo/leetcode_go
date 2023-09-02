package main

import (
	"bufio"
	"fmt"
	"os"
)

//
//var sc = bufio.NewScanner(os.Stdin)
//
//func InputRowRetInt() string {
//	sc.Scan()
//	return sc.Text()
//}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	type pair struct {
		x, y int
	}
	ans := pair{0, 0}
	for i := 0; i < len(s); i++ {
		el := s[i]
		if el == 'U' {
			ans.y += 1
		} else if el == 'D' {
			ans.y -= 1
		} else if el == 'L' {
			ans.x -= 1
		} else {
			ans.x += 1
		}
	}
	fmt.Printf("%d %d", ans.x, ans.y)
}

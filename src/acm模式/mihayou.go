package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func input() []string {
	sc.Scan()
	res := strings.Split(sc.Text(), " ")
	return res
}

func main() {
	var n, t int
	res := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		s0 := input()
		s1 := input()
		tmp := 0
		for j := 0; j < len(s0); j++ {
			if s0[j] == s1[j] {
				tmp++
			} else {
				tmp--
			}
			if tmp < 0 {
				break
			}
		}
		if tmp >= 0 {
			res++
		}
	}
	fmt.Println(res)
}

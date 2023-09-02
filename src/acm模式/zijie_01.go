package main

import (
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	s := sc.Text()
	m := map[byte]bool{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		for m[s[i]] {
			start
		}
		m[s[i]] = true
		end++
	}

}

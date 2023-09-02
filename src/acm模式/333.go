package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func InputRowRetInt() string {

	sc.Scan()

	return sc.Text()
}

func main() {
	var n int
	fmt.Scan(&n)
	s := InputRowRetInt()
	ans := n / 2
	if n%2 == 1 {
		for i := n / 2; i < n-1; i++ {
			tmp := 0
			tmp += abs(n/2 - i)
			x, y, c := i-1, i+1, 0
			for ; c < n/2; c++ {
				if s[x] != s[y] {
					tmp++
				}
				y++
				x--
				if y == n {
					y = 0
				}
			}
			ans = min(ans, tmp)
		}
	} else {
		for i := n / 2; i < n; i++ {
			tmp := 0
			tmp += abs(n/2 - i)
			x, y, c := i-1, i, 0
			for ; c < n/2; c++ {
				if s[x] != s[y] {
					tmp++
				}
				y++
				x--
				if y == n {
					y = 0
				}
			}
			ans = min(ans, tmp)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

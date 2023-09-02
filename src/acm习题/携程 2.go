package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	if n == 1 {
		fmt.Println(0)
	}

	helper := func(a, b int) int {
		if b > a {
			a, b = b, a
		}
		for a != b {
			a /= 2
			if b > a {
				a, b = b, a
			}
		}
		return b
	}

	var xVal = float64(helper(s[0], s[1]))
	for i := 2; i < n; i++ {
		xVal = float64(helper(s[i], int(xVal)))
	}
	ans := 0
	for _, el := range s {
		ans += int(math.Log2(float64(el) / xVal))
	}
	fmt.Println(ans)
}

// 最大公约数(a分母，b分子，且分母一定大于分子)
//func gcd(a, b int) int {
//	// 如果分子大于分母，需要交换顺序
//	if a < b {
//		a, b = b, a
//	}
//	for b != 0 {
//		a, b = b, a%b
//	}
//	return a
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

import "fmt"

// 暴力
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for i := 1; i <= n; i++ {
		res *= x
	}
	return res
}

// 数学  O(lgn)  以2为底n的对数
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for n > 1 {
		if n%2 == 1 {
			res *= x
			n -= 1
		}
		n /= 2
		x *= x
	}
	res *= x
	return res
}

func main() {
	fmt.Println(myPow2(2.0, 10))
}

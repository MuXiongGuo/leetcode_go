package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ans := make([]int, 10)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	if n == 1 {
		ans[s[0]]++
		for _, el := range ans {
			fmt.Printf("%d ", el%int(1e9+7))
		}
		return
	}
	helper(s, ans, n-2, s[n-1])
	for _, el := range ans {
		fmt.Printf("%d ", el%int(1e9+7))
	}
}

func helper(s, ans []int, index, c int) {
	if index < 0 {
		ans[c]++
		return
	}
	x := s[index]
	helper(s, ans, index-1, (x+c)%10)
	helper(s, ans, index-1, (x*c)%10)
}

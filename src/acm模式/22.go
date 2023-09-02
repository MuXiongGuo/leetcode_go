package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, ans int
	fmt.Scan(n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	ans = sum(s)
	if n*n >= ans {
		fmt.Println(1)
		return
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	ans -= (n + 1) * n / 2
	fmt.Println(ans)
}

func sum(s []int) int {
	ans := 0
	for _, el := range s {
		ans += el
	}
	return ans
}

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	var ans []int
	var helper func(i, val int)
	tmp := map[int]bool{}

	helper = func(i, val int) {
		if i == 3 {
			ans = append(ans, val)
			return
		}
		for j := 0; j < n; j++ {
			if tmp[j] {
				continue
			}
			tmp[j] = true
			helper(i+1, val+s[j]*int(math.Pow(10, float64(i))))
			tmp[j] = false
		}
	}
	helper(0, 0)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	for _, el := range ans {
		fmt.Printf("%d ", el)
	}
}

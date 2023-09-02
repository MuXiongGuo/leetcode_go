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
	// special
	if n == 1 {
		ans[s[0]]++
		for _, el := range ans {
			fmt.Printf("%d ", el)
		}
		return
	}
	// start
	ans[(s[n-1]*s[n-2])%10]++
	ans[(s[n-1]+s[n-2])%10]++
	tmp := make([]int, 10)
	for i := n - 3; i >= 0; i-- {
		el := s[i]
		for val := 0; val < 10; val++ {
			if ans[val] == 0 {
				continue
			}
			x1, x2 := (val*el)%10, (val+el)%10
			tmp[x1] += ans[val]
			tmp[x2] += ans[val]
			tmp[x1] = tmp[x1] % int(1e9+7)
			tmp[x2] = tmp[x2] % int(1e9+7)
		}
		for j := 0; j < 10; j++ {
			ans[j] = tmp[j]
			tmp[j] = 0
		}
	}

	for _, el := range ans {
		fmt.Printf("%d ", el%int(1e9+7))
	}
}

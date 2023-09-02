package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	ans := 0
	m := map[string]int{}
	wordGet := map[string]bool{}
	for _, el := range s {
		m[el]++
		if m[el] > ans {
			if !wordGet[el] {
				ans++
				wordGet[el] = true
			}
		}
	}
	fmt.Println(ans)
}

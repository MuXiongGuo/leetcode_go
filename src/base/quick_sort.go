package main

import "fmt"

func partition(l, r int, s []int) int {
	x := s[r]
	i := l - 1
	for j := l; j < r; j++ {
		if s[j] <= x {
			i += 1
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i+1], s[r] = s[r], s[i+1]
	return i + 1
}

func QuickSort(l, r int, s []int) {
	if l < r {
		q := partition(l, r, s)
		QuickSort(l, q-1, s)
		QuickSort(q+1, r, s)
	}
}

func main() {
	s := []int{-4, 4, 21, 6, 2, 55, 15, 23, 88, 16, -8, 60}
	QuickSort(0, len(s)-1, s)
	fmt.Println(s)
}

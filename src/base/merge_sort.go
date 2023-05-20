package main

import (
	"fmt"
	"math"
)

func merge(s []int, l, q, r int) {
	left := make([]int, q-l+2)
	right := make([]int, r-q+1)
	copy(left, s[l:q+1])
	copy(right, s[q+1:])
	left[len(left)-1] = math.MaxInt64
	right[len(right)-1] = math.MaxInt64
	i, j := 0, 0
	for k := l; k <= r; k++ {
		if left[i] <= right[j] {
			s[k] = left[i]
			i++
		} else {
			s[k] = right[j]
			j++
		}
	}
}

func MergeSort(s []int, l, r int) {
	if l < r {
		q := (l + r) / 2
		MergeSort(s, l, q)
		MergeSort(s, q+1, r)
		merge(s, l, q, r)
	}
}

func main() {
	s := []int{-4, 4, 21, 6, 2, 55, 15, 23, 88, 16, -8, 60}
	MergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}

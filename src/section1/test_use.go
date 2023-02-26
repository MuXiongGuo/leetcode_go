package main

import "sort"

func main() {
	nums := []int{2, 3, 4, 6, 7, 8, 9}
	b := sort.SearchInts(nums, 1)
	println(b)
}

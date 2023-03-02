package main

import "sort"

func search(nums []int, target int) int {
	i := sort.SearchInts(nums, target)
	n := len(nums)
	ans := 0
	for j := i; j < n; j++ {
		if nums[j] == target {
			ans++
		} else {
			break
		}
	}
	return ans
}

// 官方 太潇洒了
func search2(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}

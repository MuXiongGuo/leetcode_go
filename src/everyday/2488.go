package main

func countSubarrays(nums []int, k int) int {
	var mid int
	left, right := make([]int), []int{}
	for i := range nums {
		if nums[i] == k {
			mid = i
			break
		}
	}

}

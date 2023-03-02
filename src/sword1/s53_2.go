package main

func missingNumber(nums []int) int {
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

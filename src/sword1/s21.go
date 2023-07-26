package main

func exchange(nums []int) []int {
	j := 0
	for i, el := range nums {
		if el%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

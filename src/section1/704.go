package main

func search(nums []int, target int) int {
	//var ans int
	left, right := 0, len(nums)-1
	for {
		if left > right {
			return -1
		}
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

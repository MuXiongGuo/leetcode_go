package main

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	midIndex, minVal := 0, int(1e8+1)
	for i, v := range nums {
		if v*v < minVal {
			minVal, midIndex = v*v, i
		}
	}
	left, right := midIndex-1, midIndex+1
	ans[0] = nums[midIndex] * nums[midIndex]
	for i := 1; i < len(nums); i++ {
		if left < 0 {
			ans[i] = nums[right] * nums[right]
			right++
		} else if right >= len(nums) {
			ans[i] = nums[left] * nums[left]
			left--
		} else if nums[left]*nums[left] < nums[right]*nums[right] {
			ans[i] = nums[left] * nums[left]
			left--
		} else {
			ans[i] = nums[right] * nums[right]
			right++
		}
	}
	return ans
}

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
		if left < 0 || (right < len(nums) && nums[left]*nums[left] > nums[right]*nums[right]) {
			ans[i] = nums[right] * nums[right]
			right++
		} else {
			ans[i] = nums[left] * nums[left]
			left--
		}
	}
	return ans
}

// 官方 充分利用双指针，这次从两边开始，不断向中间，太聪明了吧，细节是index倒着
func sortedSquares2(nums []int) []int {
	ans := make([]int, len(nums))
	left, right, i := 0, len(nums)-1, len(nums)-1
	for left <= right {
		if nums[left]*nums[left] <= nums[right]*nums[right] {
			ans[i] = nums[right] * nums[right]
			right--
		} else {
			ans[i] = nums[left] * nums[left]
			left++
		}
		i--
	}
	return ans
}

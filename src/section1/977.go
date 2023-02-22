package main

// O(n)时间内解决 那么O(2n)也是O(n)
func sortedSquares(nums []int) []int {
	midIndex, minVal := 0, int(10e8+1)
	for index, val := range nums {
		if val*val < minVal {
			midIndex = index
			minVal = val * val
		}
	}
	var ans []int
	left, right := midIndex-1, midIndex+1
	ans[0] = minVal
	for left < 0 && right >= len(nums) {
		if left < 0 {
			append(ans, nums[right]*nums[right])
			right++
			continue
		}
		if right >= len(nums) {
			ans[startIndex] = nums[left] * nums[left]
			startIndex++
			left--
			continue
		}
		if nums[left]*nums[left] > nums[right]*nums[right] {
			ans[startIndex] = nums[left] * nums[left]
			startIndex++
			left--
		} else {
			ans[startIndex] = nums[right] * nums[right]
			startIndex++
			right++
		}
	}
	return ans
}

package main

func maxSubArray(nums []int) int {
	ans := nums[0]
	curMax := 0
	for _, v := range nums {
		curMax += v
		ans = max(ans, curMax)
		if curMax <= 0 {
			curMax = 0
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

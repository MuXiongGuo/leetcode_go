package main

// 本质是dp 但是不太好理解
// 不要用滑动窗口 那种统计多少个 或者 满足XXX条件的最短子数组才使用滑动窗口
func maxAbsoluteSum(nums []int) int {
	maxVal, minVal := 0, 0
	maxAns, minAns := 0, 0
	for _, el := range nums {
		maxVal += el
		minVal += el
		maxAns = max(maxAns, maxVal)
		minAns = min(minAns, minVal)
		maxVal = max(0, maxVal)
		minVal = min(0, minVal)
	}
	return max(maxAns, -minAns)
}

// 方法二：前缀和
// 由于子数组的和 等于 两个前缀和的差，那么取前缀和中的最大值与最小值，它俩的差就是答案。
// 如果最大值在最小值右边，那么算的是最大子数组和。
// 如果最大值在最小值左边，那么算的是最小子数组和的绝对值（相反数）。
func maxAbsoluteSum(nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return maxSlice(preSum) - minSlice(preSum)
}
func maxSlice(s []int) int {
	ans := s[0]
	for i := 1; i < len(s); i++ {
		if ans < s[i] {
			ans = s[i]
		}
	}
	return ans
}

func minSlice(s []int) int {
	ans := s[0]
	for i := 1; i < len(s); i++ {
		if ans > s[i] {
			ans = s[i]
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

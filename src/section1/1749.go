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

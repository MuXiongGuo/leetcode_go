package main

import (
	"math"
)

// O(n) 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	// math.MaxInt32 新东西
	ans, left, right, curVal := int(1e5+1), 0, 0, 0
	for right < len(nums) {
		curVal += nums[right]
		for curVal >= target {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			curVal -= nums[left]
			left++
		}
		right++
	}
	if ans == int(1e5+1) {
		return 0
	}
	return ans
}

// O(n*lgn) 前缀和+二分查找
// 本质上是对暴力枚举的优化
func minSubArrayLen2(target int, nums []int) int {
	n, ans := len(nums), math.MinInt32
	preSums := make([]int, n)
	for i, _ := range nums {
		if i == 0 {
			preSums[i] = nums[i]
		} else {
			preSums[i] = preSums[i-1] + nums[i]
		}
	}

}

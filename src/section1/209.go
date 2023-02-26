package main

import (
	"math"
	"sort"
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
	n, ans := len(nums), math.MaxInt32
	preSums := make([]int, n)
	for i, _ := range nums {
		if i == 0 {
			preSums[i] = nums[i]
		} else {
			preSums[i] = preSums[i-1] + nums[i]
		}
	}
	for i := 0; i <= n-1; i++ {
		if i == 0 {
			bound := sort.SearchInts(preSums, target)
			if bound == n {
				continue
			}
			ans = min(ans, bound+1)
		} else {
			bound := sort.SearchInts(preSums, target+preSums[i-1])
			if bound == n {
				continue
			}
			ans = min(ans, bound-i+1)
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func lowerBound(nums []int, target int) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {
	println(minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3}))
}

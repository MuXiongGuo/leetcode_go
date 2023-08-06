package main

// 环行的话 我们可以偷0号 与 不偷0号两种 比较一下
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	var case1, case2 int
	// case1
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < n-1; i++ {
		tmp := second
		second = max(second, nums[i]+first)
		first = tmp
	}
	case1 = second
	// case2
	first, second = nums[1], max(nums[1], nums[2])
	for i := 3; i < n; i++ {
		tmp := second
		second = max(second, nums[i]+first)
		first = tmp
	}
	case2 = second
	return max(case1, case2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

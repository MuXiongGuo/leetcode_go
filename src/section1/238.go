package main

// 1.左右乘积列表 线性空间
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := range nums {
		left[i] *= right[i]
	}
	return left
}

// 2.优化 常数空间
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	s := 1
	for i := n - 2; i >= 0; i-- {
		s *= nums[i+1]
		left[i] *= s
	}
	return left
}

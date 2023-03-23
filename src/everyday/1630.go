package main

import "sort"

// 数学
// 设b>a
// 试图O(n^2+m) 做出来  其实直接暴力+小优化O(nm) 做出来 时间复杂度差不多呢
// 失败的!!!
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n, m := len(nums), len(l)
	lookUp := make([][]bool, n)
	for i := range lookUp {
		lookUp[i] = make([]bool, n)
	}
	for i := 1; i < n; i++ {
		a, b := nums[i-1], nums[i]
		if a > b {
			a, b = b, a
		}
		s := a + b
		lookUp[i-1][i] = true
		for j := i - 2; j >= 0; j-- {
			s += nums[j]
			a = min(a, nums[j])
			b = max(b, nums[j])
			if s == (b+a)*(i-j+1)/2 {
				lookUp[j][i] = true
			}
		}
	}
	ans := make([]bool, m)
	for i := range ans {
		ans[i] = lookUp[l[i]][r[i]]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力+小优化
// 线性时间 用哈希
func isArithmeticSubarrays(nums []int) bool {
	min, max, n := nums[0], nums[0], len(nums)
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if max == min {
		return true
	}
	d := (max - min) / (n - 1)
	if (max - min) != d*(n-1) {
		return false
	}
	m := map[int]struct{}{}
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for i := min; i <= max; i += d {
		if _, ok := m[i]; !ok {
			return false
		}
	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	m := len(l)
	ans := make([]bool, m)
	for i := range l {
		ans[i] = isArithmeticSubarrays(nums[l[i] : r[i]+1])
	}
	return ans
}

// 纯暴力
func isArithmeticSubarrays(nums []int) bool {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	k := arr[1] - arr[0]
	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != k {
			return false
		}
	}
	return true
}
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	m := len(l)
	ans := make([]bool, m)
	for i := range l {
		ans[i] = isArithmeticSubarrays(nums[l[i] : r[i]+1])
	}
	return ans
}

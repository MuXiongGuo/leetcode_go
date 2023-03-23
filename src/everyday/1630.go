package main

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
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n, m := len(nums), len(l)
	lookUp := make([][]bool, n)

}

// 纯暴力
func isArithmeticSubarrays(nums []int) {

}
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n, m := len(nums), len(l)
	lookUp := make([][]bool, n)

}

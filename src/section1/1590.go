package main

// 思路是 找到一个子数组 长度最小 且满足余数c 数字p 任意数k
// 满足y-x=c+p*k
// 涉及较深的mod数学运算  最后也没写出来 但逻辑基本如此
// 还是找满足条件的子数组 只是这次要mod运算

//func minSubarray(nums []int, p int) int {
//	sumNums := 0
//	for _, el := range nums {
//		sumNums += el
//	}
//	if sumNums%p == 0 {
//		return 0
//	}
//	if sumNums < p {
//		return -1
//	}
//	sumNums %= p // 余数
//	s, n, m, ans := 0, len(nums), map[int]int{0: 0}, len(nums)
//	for i, v := range nums {
//		s += v
//		if j, ok := m[s-sumNums]; ok {
//			ans = min(ans, i-j)
//		}
//		m[s] = i
//	}
//	if ans == n {
//		return -1
//	}
//	return ans
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

// 灵神的答案
func minSubarray(nums []int, p int) int {
	x := 0
	for _, v := range nums {
		x += v
	}
	x %= p
	if x == 0 {
		return 0 // 移除空子数组（这个 if 可以不要）
	}

	n := len(nums)
	ans, s := n, 0
	// 由于下面 i 是从 0 开始的，前缀和下标就要从 -1 开始了
	last := map[int]int{s: -1}
	for i, v := range nums {
		s += v
		last[s%p] = i
		if j, ok := last[(s-x+p)%p]; ok {
			ans = min(ans, i-j)
		}
	}
	if ans < n {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

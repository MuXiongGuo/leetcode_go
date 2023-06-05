package main

// 哈希
func longestConsecutive(nums []int) int {
	m := map[int]int{}
	ans := 0
	for _, el := range nums {
		// 忽略重复数字
		if m[el] != 0 {
			continue
		}
		left := m[el-1]
		right := m[el+1]
		all := 1 + left + right
		m[el-left] = all
		m[el+right] = all
		m[el] = all
		ans = max(ans, all)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

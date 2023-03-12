package main

// 1。将字母和数字转化成+-1
// 2。前缀和可以快速解决最长子数组问题
// 3。统计时用个哈希即可
func findLongestSubarray(array []string) []string {
	m := map[int]int{}
	s := make([]int, len(array)+1)
	for i, v := range array {
		if v[0] >= '0' && v[0] <= '9' {
			s[i+1] = s[i] + 1
		} else {
			s[i+1] = s[i] - 1
		}
	}
	start, end := 0, 0
	for i, v := range s {
		if j, ok := m[v]; ok {
			if i-j > end-start {
				end, start = i, j
			}
		} else {
			m[v] = i
		}
	}
	return array[start:end]
}

// 灵茶山优化！！！
// 1。避免了if else判断
// 2。前缀和数组用变量代替
// 3。哈希表用2n+1数组代替 因为范围是-n到n
func findLongestSubarray(array []string) []string {
	begin, end, n := 0, 0, len(array)
	first := make([]int, n*2+1)
	for i := range first { // 注：去掉可以再快 1ms（需要 s 下标改从 1 开始）
		first[i] = -1
	}
	s := n
	first[s] = 0 // s[0] = 0
	for i := 1; i <= n; i++ {
		s += int(array[i-1][0])>>6&1*2 - 1
		if j := first[s]; j < 0 {
			first[s] = i
		} else if i-j > end-begin {
			begin, end = j, i
		}
	}
	return array[begin:end]
}

package main

func findMaxLength(nums []int) int {
	s, ans, m := 0, 0, map[int]int{0: -1}
	for i, v := range nums {
		s += v*2 - 1
		if j, ok := m[s]; !ok {
			m[s] = i
		} else if i-j > ans {
			ans = i - j
		}
	}
	return ans
}

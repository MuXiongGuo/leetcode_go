package main

// 好像只能移动一次 并且很城堡看起来很蠢
func captureForts(forts []int) int {
	cur, ans := 0, 0
	for i := 1; i < len(forts); i++ {
		if forts[i] == 1 || forts[i] == -1 {
			if forts[i]*forts[cur] < 0 {
				ans = max(ans, i-cur-1)
			}
			cur = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

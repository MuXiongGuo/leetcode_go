package main

// 滑动窗口+哈希 (set)
func lengthOfLongestSubstring(s string) int {
	m := map[uint8]struct{}{}
	n := len(s)
	ans := 0
	for i, j := 0, 0; i < n; i++ {
		for _, ok := m[s[i]]; ok; {
			delete(m, s[j])
			j++
			_, ok = m[s[i]]
		}
		m[s[i]] = struct{}{}
		ans = max(ans, i-j+1)
	}
	return ans
}

// 坐标的map 更优雅
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	ans := 0
	left, right := 0, 0
	for right < n {
		if i, ok := m[s[right]]; ok && i >= left {
			left = i + 1
		}
		m[s[right]] = right
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	lengthOfLongestSubstring("abcabcbb")
}

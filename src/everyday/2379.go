package main

// 滑动窗口
func minimumRecolors(blocks string, k int) int {
	cur := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			cur++
		}
	}
	ans := cur
	n := len(blocks)
	for i, j := k, 0; i < n; i, j = i+1, j+1 {
		if blocks[j] == 'W' {
			cur--
		}
		if blocks[i] == 'W' {
			cur++
		}
		if cur < ans {
			ans = cur
		}
	}
	return ans
}

// 官方
// 挺优雅
func minimumRecolors(blocks string, k int) int {
	res := k
	left, whites := 0, 0
	for right := 0; right < len(blocks); right++ {
		if blocks[right] == 'W' {
			whites++
		}
		if right < k-1 {
			continue
		}
		res = min(res, whites)
		if blocks[left] == 'W' {
			whites--
		}
		left++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

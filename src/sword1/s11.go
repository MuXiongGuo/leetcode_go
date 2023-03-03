package main

// 自己的 比较乱 有漏洞
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	// 特殊情况 没旋转
	if numbers[right] >= numbers[left] {
		return numbers[left]
	}
	// 正常
	x := numbers[0]
	ans := 5001
	for left <= right {
		mid := (right-left)/2 + left
		ans = min(ans, numbers[mid])
		if numbers[mid] < x {
			right = mid - 1
		} else if numbers[mid] > x {
			left = mid + 1
		} else {
			right--
		}
	}
	return ans
}

// 不能用固定第一个头节点 有破绽 要用两边的随时更新

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

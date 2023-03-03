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
// 其实很简单 不要想的太复杂 就是mid与右端点判断即可
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	ans := 5001
	for low <= high {
		pivot := low + (high-low)/2
		ans = min(ans, numbers[pivot])
		if numbers[pivot] < numbers[high] {
			high = pivot - 1
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 官方这个更聪明
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

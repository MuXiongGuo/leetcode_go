package main

// 线性空间
func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		x := min(left[i-1], right[i+1]) - height[i]
		if x < 0 {
			x = 0
		}
		ans += x
	}
	return ans
}

//	func main() {
//		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
//	}
//
// 左右指针 O(1)空间  判断左右 小的一方移动那一方的 利用已有的信息量来完成题目
func trap(height []int) int {
	ans := 0
	l, r := height[0], height[len(height)-1]
	i, j := 1, len(height)-2
	for i <= j {
		if l > r {
			x := min(l, r) - height[j]
			if x > 0 {
				ans += x
			}
			r = max(r, height[j])
			j--
		} else {
			x := min(l, r) - height[i]
			if x > 0 {
				ans += x
			}
			l = max(l, height[i])
			i++
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

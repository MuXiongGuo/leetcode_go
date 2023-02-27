package main

// 没必要弄两个数组，因为不会修改 只是统计而已
func movesToMakeZigzag(nums []int) int {
	evenTimes, oddTimes, n := 0, 0, len(nums)
	evenNums, oddNums := make([]int, n), nums[:]
	copy(evenNums, oddNums)
	// even
	for i := 1; i < n; i += 2 {
		if i == n-1 {
			if evenNums[i] <= evenNums[i-1] {
				evenTimes += evenNums[i-1] - evenNums[i] + 1
			}
		} else {
			if evenNums[i] <= evenNums[i-1] {
				evenTimes += evenNums[i-1] - evenNums[i] + 1
			}
			if evenNums[i] <= evenNums[i+1] {
				evenTimes += evenNums[i+1] - evenNums[i] + 1
				evenNums[i+1] = evenNums[i] - 1
			}
		}
	}

	// odd
	for i := 1; i < n; i += 2 {
		if i == n-1 {
			if oddNums[i] >= oddNums[i-1] {
				oddTimes += oddNums[i] - oddNums[i-1] + 1
			}
		} else {
			if oddNums[i] >= min(oddNums[i+1], oddNums[i-1]) {
				oddTimes += oddNums[i] - min(oddNums[i+1], oddNums[i-1]) + 1
			}
		}
	}
	if oddTimes < evenTimes {
		return oddTimes
	}
	return evenTimes
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 官方，相似代码复用，值得学习
// 奇数偶数 本质是一样的逻辑
func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	help := func(pos int) int {
		ans := 0
		for i := pos; i < n; i += 2 {
			if i == 0 {
				if n >= 2 {
					ans += max(0, nums[i]-nums[i+1]+1)
				}
			} else if i == n-1 {
				ans += max(0, nums[i]-nums[i-1]+1)
			} else {
				ans += max(0, nums[i]-min(nums[i+1], nums[i-1])+1)
			}
		}
		return ans
	}
	n1, n2 := help(0), help(1)
	if n1 < n2 {
		return n1
	}
	return n2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

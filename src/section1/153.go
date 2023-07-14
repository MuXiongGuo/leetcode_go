package main

func findMin(nums []int) int {
	n := len(nums)
	// case: n=1 以及 顺序状态
	if nums[0] <= nums[n-1] {
		return nums[0]
	}
	compareValue := nums[0]
	checkQualify := func(x int) bool {
		if x == 0 {
			return false
		}
		if nums[x-1] > nums[x] {
			return true
		}
		return false
	}

	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if checkQualify(mid) {
			return nums[mid]
		}
		if nums[mid] >= compareValue {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

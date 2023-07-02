package main

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return true
	}

	l, r := 0, n-1
	helper := func(i int) (bool, int) {
		if i == 0 {
			if nums[i+1] < nums[i] {
				return true, i
			}
			return false, i + 1
		} else if i == n-1 {
			if nums[i-1] < nums[i] {
				return true, i
			}
			return false, i - 1
		}

		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return true, i
		}
		if nums[i] < nums[i+1] {
			return false, i + 1
		}

		return false, i - 1
	}

	for l <= r {

		mid := (l + r) / 2
		if ok, x := helper(mid); ok {
			return x
		} else if x > mid {
			l = x
		} else {
			r = x
		}
	}

	return 0
}

func main() {
	findPeakElement([]int{1})
}

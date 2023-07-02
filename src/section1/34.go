package main

import "sort"

func searchRange(nums []int, target int) []int {
	ans := []int{}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	ans = append(ans, l)

	l, r = 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	ans = append(ans, r)

	if ans[0] < 0 || ans[1] >= len(nums) || nums[ans[0]] != target || nums[ans[1]] != target {
		return []int{-1, -1}
	}

	return ans
}

func searchRange(nums []int, target int) []int {
	l := sort.SearchInts(nums, target)
	r := sort.SearchInts(nums, target+1) - 1
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	return []int{l, r}

}

func main() {
	searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
}

package main

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	if nums[n-2] < nums[n-1] {
		nums[n-1], nums[n-2] = nums[n-2], nums[n-1]
		return
	}
	maxVal := nums[n-2]
	for i := n - 3; i >= 0; i-- {
		if nums[i] >= maxVal {
			maxVal = nums[i]
		} else {
			for j := n - 1; j >= 0; j-- {
				if nums[i] < nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			tmp := nums[i+1:]
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i] < tmp[j]
			})
			return
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return
}

func main() {
	nextPermutation([]int{5, 4, 7, 5, 3, 2})
}

//	func main() {
//		n := []int{1, 4, 5, 2, 3, 5, 9}
//		tmp := n[2:]
//		sort.Slice(tmp, func(i, j int) bool {
//			return tmp[i] < tmp[j]
//		})
//		fmt.Println(n)
//	}
//
// 官方 聪明！！ 其实是翻转 reverse 不是排序 因为是逆序的节约时间
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

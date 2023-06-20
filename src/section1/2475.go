package main

import "sort"

// 1. 回溯先试试
func helper(nums, q []int, start int, ans *int) {
	if len(q) == 3 {
		*ans++
		return // 不停止就成了2^n  停止是n^3
	}
	for i := start; i < len(nums); i++ {
		if find(q, nums[i]) == false {
			q = append(q, nums[i])
			helper(nums, q, i+1, ans)
			q = q[:len(q)-1]
		}
	}

}

func unequalTriplets(nums []int) int {
	var ans = 0
	q := []int{}
	helper(nums, q, 0, &ans)
	return ans
}

func find(n []int, k int) bool {
	for _, el := range n {
		if el == k {
			return true
		}
	}
	return false
}

// 2. 排序后便于统计
func unequalTriplets(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i, j := 0, 0; i < len(nums); i = j {
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		res += i * (j - i) * (len(nums) - j)
	}
	return res
}

// 3. 排序只是为了让相同的在一起 顺序无所谓的 这太数学了
func unequalTriplets(nums []int) int {
	count := map[int]int{}
	for _, x := range nums {
		count[x]++
	}
	res, n, t := 0, len(nums), 0
	for _, v := range count {
		res, t = res+t*v*(n-t-v), t+v
	}
	return res
}

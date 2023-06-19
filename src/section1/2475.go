package main

// 回溯先试试
func helper(nums, q []int, start int, ans *int) {
	if len(q) == 3 {
		*ans++
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

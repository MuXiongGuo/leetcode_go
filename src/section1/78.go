package main

// 子集问题不是for循环
// 它是每一步都有选或不选
func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	var path []int
	var helper func(i int)
	helper = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...)) // 这里应该是path的copy
			return
		}
		helper(i + 1)

		path = append(path, nums[i])
		helper(i + 1)
		path = path[:len(path)-1]
	}
	helper(0)
	return ans
}

// 列举所有状态 每个状态花费O(n)时间去构造
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func main() {
	subsets([]int{1, 2, 3})
}

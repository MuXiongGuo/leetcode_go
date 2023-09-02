package main

import (
	"fmt"
	"math/rand"
)

// O(n)时间内完成 找到第k大  其实找前k大也是O(n) 平均O(n) 算法导论分析过
// 附带如何产生随机数
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, l, r, k int) int {
	i := partition(nums, l, r)
	if k == i+1 {
		return nums[i]
	}
	if k > i+1 {
		return quickSelect(nums, i+1, r, k) // k不需要变化的!
	}
	return quickSelect(nums, l, i-1, k)
}

func partition(nums []int, l, r int) int {
	//x := (rand.Int() + l) % (r) 这是个错误的随机区间算法
	x := rand.Int()%(r-l+1) + l // 这个才是正确的
	i, j := l-1, l
	nums[x], nums[r] = nums[r], nums[x]
	for ; j < r; j++ {
		if nums[j] > nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[j] = nums[j], nums[i]
	return i
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

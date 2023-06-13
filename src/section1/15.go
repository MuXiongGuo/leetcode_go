package main

import "sort"

// 不是O(n)!!!
// 两数之和进阶
package main

import "sort"

// 不是O(n)!!!
// 两数之和进阶
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 两个优化
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		// 核心是枚举的i为左端点, 而不是中点, 如果中点会因为-1 -1 2 这种例子被忽略掉  枚举左端就可以!
		j, k := i+1, n-1
		for j < k {
			x := nums[i] + nums[j] + nums[k]
			if x > 0 {
				k--
			} else if x < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				for j < n && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > 0 && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ans
}
func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}

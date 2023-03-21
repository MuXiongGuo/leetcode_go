package main

// 1.哈希前缀和统计
func countSubarrays(nums []int, k int) int {
	mLeft, s, ans := map[int]int{}, 0, 0
	var start int
	for i, v := range nums {
		if v == k {
			start = i
			break
		}
	}
	for i := start; i >= 0; i-- {
		if nums[i] > k {
			s += 1
		} else if nums[i] < k {
			s -= 1
		}
		mLeft[s]++
	}
	s = 0
	ans += mLeft[0] + mLeft[1]
	for i := start + 1; i < len(nums); i++ {
		if nums[i] > k {
			s += 1
		} else if nums[i] < k {
			s -= 1
		}
		ans += mLeft[-s] + mLeft[-s+1]
	}
	return ans
}
func main() {
	//[3,2,1,4,5], k = 4
	countSubarrays([]int{3, 2, 1, 4, 5}, 4)
}

package main

// 注意left可能会超过right的细节即可
func numSubarrayProductLessThanK(nums []int, k int) int {
	left, ans, cur, n, cnt := 0, 0, 1, len(nums), 0
	for right := 0; right < n; right++ {
		cur *= nums[right]
		cnt++
		for cur >= k && left <= right {
			cur /= nums[left]
			left++
			cnt--
		}
		ans += cnt
	}
	return ans
}

// 灵神 完美！！！  right-left 已经是长度了！！！
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	left, ans, cur, n := 0, 0, 1, len(nums)
	for right := 0; right < n; right++ {
		cur *= nums[right]
		for cur >= k {
			cur /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

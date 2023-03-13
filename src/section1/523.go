package main

func checkSubarraySum(nums []int, k int) bool {
	s, m := 0, map[int]int{0: -1}
	for i, v := range nums {
		s += v // remainder = (remainder + num) % k  这样写不容易溢出
		if j, ok := m[s%k]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			m[s%k] = i
		}
	}
	return false
}

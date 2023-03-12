package main

// 哈希+前缀和  同样适合解决子数组问题
func subarraySum(nums []int, k int) int {
	s := 0
	m := map[int]int{}
	m[0] = 1
	ans := 0
	for _, v := range nums {
		s += v
		if i, ok := m[s-k]; ok {
			ans += i
		}
		m[s]++
	}
	return ans
}

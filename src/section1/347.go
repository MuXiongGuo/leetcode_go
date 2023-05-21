package main

import "sort"

// 就算没有order map  也可以用unorder map再排序一样的效果
// O(n*lgn) 复杂度
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, el := range nums {
		m[el]++
	}
	type pair struct {
		val   int
		count int
	}
	var s []pair
	for k, el := range m {
		s = append(s, pair{
			val:   k,
			count: el,
		})
	}
	// 从大到小排序
	sort.Slice(s, func(i, j int) bool {
		return s[i].count > s[j].count
	})
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, s[i].val)
	}
	return ans
}

package main

import (
	"sort"
)

// sort 定制操作
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var ans [][]int
	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})
	sort.Slice(items2, func(i, j int) bool {
		return items2[i][0] < items2[j][0]
	})
	n, m := len(items1), len(items2)
	i, j := 0, 0
	for i != n || j != m {
		if i == n || (j != m && items1[i][0] > items2[j][0]) {
			ans = append(ans, items2[j])
			j++
		} else if j == m || (i != n && items1[i][0] < items2[j][0]) {
			ans = append(ans, items1[i][:])
			i++
		} else {
			ans = append(ans, []int{items1[i][0], items1[i][1] + items2[j][1]})
			i++
			j++
		}
	}
	return ans
}

// 官方 哈希+排序 可惜没有红黑树 又要多一个排序
func mergeSimilarItems(item1, item2 [][]int) [][]int {
	m := map[int]int{}
	for _, v := range item1 {
		m[v[0]] += v[1]
	}
	for _, v := range item2 {
		m[v[0]] += v[1]
	}
	var ans [][]int
	for a, b := range m {
		ans = append(ans, []int{a, b})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

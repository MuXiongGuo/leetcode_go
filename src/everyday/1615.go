package main

import "sort"

type pair struct {
	index int
	count int
}

type pairSlice []pair

func (p pairSlice) Less(i, j int) bool {
	return p[i].count > p[j].count
}
func (p pairSlice) Len() int {
	return len(p)
}
func (p pairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 两种方法实现排序pair问题
// 1.可以直接暴力 复杂度O(n^2)
// 2.不sort 通过集合判断 很优秀 具体解析 以及时间复杂度分析可以看 官方!!!
// O(m+n)  找到第一大集合并统计个数 为1 则枚举第2大里面 时间不会超过n
// 若较多x  则判断x(x-1)/2 与 m关系 可直接判断
// 就算是小于 也可以继续枚举 时间不会超过道路总数m 所以还是m+n  m+n包括统计邻接表的时间
// 空间复杂分析: 邻接表O(n+m) 若用数组O(n^2)

func maximalNetworkRank(n int, roads [][]int) int {
	adList := make([][]int, n)
	for i := range adList {
		adList[i] = make([]int, n)
	}
	pCount := make([]pair, n)
	for i := range pCount {
		pCount[i].index = i
	}
	for i := 0; i < len(roads); i++ {
		adList[roads[i][0]][roads[i][1]]++
		adList[roads[i][1]][roads[i][0]]++
		pCount[roads[i][0]].count++
		pCount[roads[i][1]].count++
	}
	sort.Slice(pCount, func(i, j int) bool {
		return pCount[i].count > pCount[j].count
	})
	max1, max2 := pCount[0].count, pCount[1].count
	index1 := pCount[0].index
	// 分类讨论
	if max1 != max2 {
		for i := 1; i < n; i++ {
			if pCount[i].count == max2 && adList[pCount[i].index][index1] == 0 {
				return max1 + max2
			}
		}
		return max1 + max2 - 1
	} else {
		index2 := 1
		for i := 1; i < n; i++ {
			if pCount[i].count == max1 {
				index2 = i
			}
		}
		for i := 0; i <= index2; i++ {
			for j := i + 1; j <= index2; j++ {
				if adList[pCount[i].index][pCount[j].index] == 0 {
					return max1 + max2
				}
			}
		}
		return max1 + max2 - 1
	}
}
func main() {
	//a := pairSlice{{1, 2}, {2, 3}, {3, 4}}
	////sort.Sort(a)
	//sort.Slice(a, func(i, j int) bool { return a[i].count > a[j].count })
	//println(a)

	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}
	maximalNetworkRank(4, roads)
}

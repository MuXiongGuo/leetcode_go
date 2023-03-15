package main

import "sort"

type pair struct {
	index int
	count int
}

type pairSlice []pair

func (p []pair) Less(i, j int) bool {
	return p[i].count > p[j].count
}
func (p pairSlice) Len() int {
	return len(p)
}
func (p pairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//	func maximalNetworkRank(n int, roads [][]int) int {
//		adList := make([][]int, n)
//		adCount := make([]int, n)
//
// }
func main() {
	a := pairSlice{{1, 2}, {2, 3}, {3, 4}}
	sort.Sort(a)
	println(a)
}

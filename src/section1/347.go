package main

import (
	"container/heap"
	"sort"
)

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

// 使用堆 则可以O(n*lgk)完成 k为前k高频
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, v := range nums {
		occurrences[v]++
	}
	h := &hp{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = heap.Pop(h).([2]int)[0]
	}
	return ret
	//h := &hp{sort.IntSlice{3, 6, 4, 1}}
	//heap.Init(h)
	//heap.Push(h, 99)
	//n := h.Len()
	//for i := 0; i < n; i++ {
	//	fmt.Println(heap.Pop(h).(int))
	//}
}

//type hp struct {
//	sort.IntSlice
//}
//
//func (s *hp) Less(i, j int) bool {
//	return s.IntSlice[i] > s.IntSlice[j]
//}
//
//func (s *hp) Push(v interface{}) {
//	s.IntSlice = append(s.IntSlice, v.(int))
//}
//
//func (s *hp) Pop() interface{} {
//	ans := s.IntSlice[len(s.IntSlice)-1]
//	s.IntSlice = s.IntSlice[:len(s.IntSlice)-1]
//	return ans
//}

type hp [][2]int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *hp) Pop() interface{} {
	ans := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ans
}

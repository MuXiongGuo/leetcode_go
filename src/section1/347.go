package main

import (
	"container/heap"
	"math/rand"
	"sort"
	"time"
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

// 前k高频模板
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	values := [][]int{}
	for key, value := range occurrences {
		values = append(values, []int{key, value})
	}
	ret := make([]int, k)
	qsort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		qsort(values, start, index-1, ret, retIndex, k)
	} else {
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}

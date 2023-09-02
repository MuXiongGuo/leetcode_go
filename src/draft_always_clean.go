package main

import (
	"math/rand"
	"time"
)

// 自己手写一下 TopK模板
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, val := range nums {
		m[val]++
	}
	var q [][]int
	for k, v := range m {
		q = append(q, []int{k, v})
	}
	var res []int
	qSort(q, &res, 0, len(q)-1, k, 0)
	return res
}

func qSort(nums [][]int, res *[]int, l, r, k, curResIdx int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Int()%(r-l+1) + l
	nums[x], nums[r] = nums[r], nums[x]
	i, j := l-1, l
	for ; j < r; j++ {
		if nums[j][1] > nums[r][1] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[j] = nums[j], nums[i]
	if k < i+1-l {
		qSort(nums, res, l, i-1, k, curResIdx)
	} else {
		for _, el := range nums[l : i+1] {
			*res = append(*res, el[0])
		}
		if i+1-l == k { // 已经收集满足了
			return
		}
		qSort(nums, res, i+1, r, k-i-1+l, curResIdx)
	}
}
func main() {
	topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
}

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, last, buy int
	fmt.Scan(&n)
	var ans []int
	t := &tps{}
	heap.Init(t)
	lastS, buyS := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lastS[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&buyS[i])
	}
	for i := 0; i < n; i++ {
		last = lastS[i]
		buy = buyS[i]
		heap.Push(t, tuple{
			month: i + 1,
			last:  last,
			val:   i + 1 + last,
		})
		for t.Len() != 0 && buy != 0 {
			el := heap.Pop(t).(tuple)
			if i+1 <= el.last+el.month {
				ans = append(ans, el.month)
				buy--
			}
		}
	}
	fmt.Println(len(ans))
	for _, el := range ans {
		fmt.Printf("%d ", el)
	}
}

type tuple struct {
	month int
	last  int
	val   int
}

type tps []tuple

func (t tps) Len() int { return len(t) }
func (t tps) Less(i, j int) bool {
	if t[i].val < t[j].val {
		return true
	} else if t[i].val > t[j].val {
		return false
	} else {
		return t[i].month < t[j].month
	}
}
func (t tps) Swap(i, j int)       { t[i], t[j] = t[j], t[i] }
func (t *tps) Push(x interface{}) { *t = append(*t, x.(tuple)) }
func (t *tps) Pop() interface{} {
	ans := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return ans
}

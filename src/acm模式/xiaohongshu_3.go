package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// input
	var n, m, k int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	aValue := make([]int, n+1)
	hCost := make([]int, n+1)
	adList := make([][]pair, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&aValue[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&hCost[i])
	}
	for i := 1; i <= m; i++ {
		var u, v, w int
		fmt.Scan(&u)
		fmt.Scan(&v)
		fmt.Scan(&w)
		adList[u] = append(adList[u], pair{v: v, w: w})
	}
	used := map[int]bool{}
	// start

}

func dijkstra(adList [][]pair, start int) []int {
	dist := make([]int, len(adList))
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.v
		if dist[x] < p.w {
			continue
		}
		for _ , el := range adList[x] {
			y := el.v
			if d := dist[x]+el.v
		}
	}
}

type pair struct{ v, w int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].v < h[j].v }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { old := *h; *h, v = old[:len(old)-1], old[len(old)-1]; return }

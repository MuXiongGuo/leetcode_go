package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type pair struct {
	val   int
	index int
}

type hp []pair

func mergeKLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}
	p := dummyNode
	h := &hp{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, pair{val: lists[i].Val, index: i})
			lists[i] = lists[i].Next
		}
	}
	for h.Len() != 0 {
		cur := heap.Pop(h).(pair)
		p.Next = &ListNode{Val: cur.val}
		if lists[cur.index] != nil {
			heap.Push(h, pair{val: lists[cur.index].Val, index: cur.index})
			lists[cur.index] = lists[cur.index].Next
		}
		p = p.Next
	}
	return dummyNode.Next
}

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	return (*h)[i].val < (*h)[j].val
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *hp) Pop() interface{} {
	old := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return old
}

func main() {
	// [[1,4,5],[1,3,4],[2,6]]
	mergeKLists(
		[]*ListNode{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	)
}

package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	l := []int{}
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	dummyNode := &ListNode{}
	p := dummyNode
	for i := 0; i < len(l); i++ {
		p.Next = &ListNode{Val: l[i]}
		p = p.Next
	}
	return dummyNode.Next
}

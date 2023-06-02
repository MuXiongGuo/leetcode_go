package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	p := dummyNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			p.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		p = p.Next
	}
	for list1 != nil {
		p.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		p = p.Next
	}
	for list2 != nil {
		p.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		p = p.Next
	}
	return dummyNode.Next
}

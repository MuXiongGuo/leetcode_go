package main

// 修改了这两个链表 懒得再复制新的
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyNode *ListNode = &ListNode{
		Val:  0,
		Next: nil,
	}
	head := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
			head = head.Next
		} else {
			head.Next = l2
			l2 = l2.Next
			head = head.Next
		}
	}
	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummyNode.Next
}

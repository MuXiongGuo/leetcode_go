package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	p0 := dummy
	for p0.Next != nil && p0.Next.Next != nil {
		p1 := p0.Next
		p2 := p1.Next
		p1.Next = p2.Next
		p2.Next = p1
		p0.Next = p2
		// 前进到下一组
		p0 = p1
	}
	return dummy.Next
}

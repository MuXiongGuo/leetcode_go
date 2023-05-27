package main

// 按照灵神 一步一步来
func reverseKGroup(head *ListNode, k int) *ListNode {
	listLen := 0
	p := head
	for p != nil {
		listLen++
		p = p.Next
	}

	dummyNode := &ListNode{Next: head}
	p0 := dummyNode
	for listLen >= k {
		listLen -= k
		var cur, pre *ListNode = p0.Next, nil
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		p0Save := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		// p0 随之移动即可
		p0 = p0Save
	}
	return dummyNode.Next
}

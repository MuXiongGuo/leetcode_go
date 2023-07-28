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



// 时隔多日 我自己写的 一次性通过
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	dummyNode := &ListNode{Next: head}
	n := 0
	for p := dummyNode.Next; p != nil; p = p.Next {
		n++
	}
	start := dummyNode
	for n >= k {
		n -= k
		q0, q1 := start.Next, start.Next.Next
		var q2 *ListNode
		for i := 1; i < k; i++ {
			q2 = q1.Next
			q1.Next = q0
			q0 = q1
			q1 = q2
		}
		// 此时 q0为尾部   q1=q2 为q0下一个节点
		q2 = start.Next
		start.Next.Next = q1
		start.Next = q0
		start = q2
	}
	return dummyNode.Next

}
func main() {
	reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
}

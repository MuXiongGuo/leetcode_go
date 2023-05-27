package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	l, r := head, head
	var pre, cur *ListNode
	cur = r.Next
	for i := 1; i < left; i++ {
		pre = l
		l = l.Next
	}
	for i := 1; i < right; i++ {
		r = r.Next
		cur = r.Next
	}
	if pre != nil {
		pre.Next = r
	}

	var pPre, pNext, pCur *ListNode
	pCur, pPre = l, cur
	for pCur != cur {
		pNext = pCur.Next
		pCur.Next = pPre
		pPre = pCur
		pCur = pNext
	}
	if pre == nil {
		return r
	}
	return head
}

// 使用dummyNode避免特殊情况
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	p0 := dummyNode
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	var cur, pre *ListNode = p0.Next, nil
	for i := 0; i <= right-left; i++ {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	p0.Next.Next = cur
	p0.Next = pre
	return dummyNode.Next
}

func main() {
	reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, 4)
}

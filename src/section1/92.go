package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	l, r := head, head
	var pre, next *ListNode
	for i := 1; i < left; i++ {
		pre = l
		l = l.Next
	}
	for i := 1; i < right; i++ {
		r = r.Next
		next = r.Next
	}
	if pre != nil {
		pre.Next = r
	}

	cur, p := l, next
	var n *ListNode
	end := r.Next
	for cur != end {
		n = cur.Next
		cur.Next = p
		p = cur
		cur = n
	}
	if pre == nil {
		return r
	}
	return pre
}
func main() {
	reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, 4)
}

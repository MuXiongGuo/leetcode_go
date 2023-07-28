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

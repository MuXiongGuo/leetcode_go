package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 虚拟头节点 迭代法
func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head} // 不需要make
	for p := dummyNode; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummyNode.Next
}

// 递归 很巧妙的方法
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements2(head.Next, val)
	}
	head.Next = removeElements2(head.Next, val)
	return head
}

// 递归 很巧妙的写法
func removeElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements3(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

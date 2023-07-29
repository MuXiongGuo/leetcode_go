package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head, Val: 101}
	pre, cur := dummy, dummy.Next
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur = pre.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 更简洁
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil { // 不需要走到最后，在最后一个的前一个就可以判断它是否符合，而且删除也方便！！！
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

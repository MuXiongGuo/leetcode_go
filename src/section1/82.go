package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 删除逻辑和不删除逻辑分开  很聪明！
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 101, Next: head}
	pre, cur := dummy, dummy.Next
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			next := cur.Next
			for next != nil && next.Val == cur.Val {
				next = next.Next
			}
			pre.Next = next
			cur = next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

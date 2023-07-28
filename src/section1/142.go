package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针相遇时的节点 与 头节点到 入口(ans) 是整数倍关系，数学！
// 故开始不断循环即可
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 哈希或者快慢指针
func hasCycle(head *ListNode) bool {
	quick, slow := head, head
	for quick != nil {
		slow = slow.Next
		quick = quick.Next
		if quick == nil {
			return false
		}
		quick = quick.Next
		if slow == quick {
			return true
		}
	}
	return false
}

// 哈希
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 官方
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

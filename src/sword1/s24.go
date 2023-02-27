package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 完美递归 注意小坑可能有空链表进入
// 用迭代也可以 只要事先存储下Next节点
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ans
}

// 迭代
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for p := head; p != nil; {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}

func main() {
	reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
}

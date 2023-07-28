package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先用容器装起来  有个小坑 nil的尾部要添加好！
func reorderList(head *ListNode) {
	var deque []*ListNode
	p := head
	for p != nil {
		deque = append(deque, p)
		p = p.Next
	}
	dummyNode := &ListNode{}
	p0 := dummyNode
	i, j := 0, len(deque)-1
	for i <= j {
		p0.Next = deque[i]
		i++
		p0 = p0.Next
		p0.Next = nil
		if i > j {
			break
		}
		p0.Next = deque[j]
		j--
		p0 = p0.Next
		p0.Next = nil
	}
	// return dummyNode.Next
}

// 灵神与官方 O(1)空间复杂度
// 1. 找到中点 2. 反转后半部分 3. 合并两个链表
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func main() {
	reorderList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}})
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func helper(head *ListNode, ans *[]int) {
	if head == nil {
		return
	}
	helper(head.Next, ans)
	ans = append(ans, head.Val)
}
func reversePrint(head *ListNode) []int {
	var ans []int
	helper(head, &ans)
	return ans
}

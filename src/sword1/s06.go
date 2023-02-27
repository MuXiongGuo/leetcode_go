package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 递归做法, 很经典的用递归来删除 倒序
// 重点是学学切片的底层 以及切片的指针
func helper(head *ListNode, ans *[]int) {
	if head == nil {
		return
	}
	helper(head.Next, ans)
	*ans = append(*ans, head.Val)
}
func reversePrint(head *ListNode) []int {
	var ans []int
	helper(head, &ans)
	return ans
}

package main

// 脑筋急转弯，把前一个copy过来，再循环往复 最后删除一个
// 其实不用循环！！！
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		root = queue[0]
		queue = queue[1:]
		ans = append(ans, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}
	return ans
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 跟题目100 一个套路 完美判断法！
func isSymmetric(root *TreeNode) bool {
	var helper func(p1, p2 *TreeNode) bool
	helper = func(p1, p2 *TreeNode) bool {
		if p1 == nil || p2 == nil {
			return p1 == p2
		}
		return p1.Val == p2.Val && helper(p1.Left, p2.Right) && helper(p1.Right, p2.Left)
	}
	return helper(root.Left, root.Right)
}

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	var helper func(p *TreeNode) (bool, int)
	helper = func(p *TreeNode) (bool, int) {
		if p == nil {
			return true, 0
		}
		flag1, left := helper(p.Left)
		flag2, right := helper(p.Right)
		flag := flag1 && flag2 && (abs(left-right) <= 1)
		return flag, max(left, right) + 1
	}
	ans, _ := helper(root)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 灵神！ 更加简介和优雅 没必要增加 bool 选项，利用高度为-1作为不合格的标记，并且提前结束
func isBalanced(root *TreeNode) bool {
	var helper func(p *TreeNode) int
	helper = func(p *TreeNode) int {
		if p == nil {
			return 0
		}
		left := helper(p.Left)
		if left == -1 {
			return -1
		}
		right := helper(p.Right)
		if right == -1 || abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	return helper(root) != -1
}

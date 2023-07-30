package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 判断二叉搜索
func isValidBST(root *TreeNode) bool {
	var helper = func(p *TreeNode) (int, int) {
		if p == nil {
			return 0, 0
		}
	}

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

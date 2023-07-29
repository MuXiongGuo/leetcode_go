package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		for n := len(q); n > 0; n-- {
			p := q[0]
			if p.Left != nil {
				q = append(q, p.Left)
			}
			if p.Right != nil {
				q = append(q, p.Right)
			}
			if n == 1 {
				ans = append(ans, q[0].Val)
			}
			q = q[1:]
		}
	}
	return ans
}

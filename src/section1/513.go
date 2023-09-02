package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	ans := [][]int{}
	q := []*TreeNode{root}
	for len(q) != 0 {
		n := len(q)
		tmp := []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
		ans = append(ans, tmp)
	}
	return ans[len(ans)-1][0]
}

// 灵神 直接改成先右儿子再左儿子 这样最后一个就是目标啦!
func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	var node *TreeNode
	for len(q) != 0 {
		node = q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}

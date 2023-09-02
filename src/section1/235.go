package main

// 搜索的话没必要用模板了
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if root.Val < max(p.Val, q.Val) && root.Val > min(p.Val, q.Val) {
		return root
	}
	if root.Val > max(p.Val, q.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
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

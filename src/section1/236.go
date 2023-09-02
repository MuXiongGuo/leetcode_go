package main

// 分类讨论
// 1. 节点为空则返回此节点
// 2. 节点为p或q直接返回此节点
// 3. 节点左右子树为 p q 返回这个节点此为公共祖先
// 4. 只有左子树或右子树中有 p q 则返回此节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

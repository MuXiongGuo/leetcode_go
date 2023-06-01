package main

// dfs
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	//left := mirrorTree(root.Left)
	//right := mirrorTree(root.Right)
	//root.Left, root.Right = right, left
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

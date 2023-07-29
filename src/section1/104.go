package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var dfs func(p *TreeNode) int
	dfs = func(p *TreeNode) int {
		if p == nil {
			return 0
		}
		left := dfs(p.Left)
		right := dfs(p.Right)
		return max(left, right) + 1
	}
	return dfs(root)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 更加优雅！
//package main
//
///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//func maxDepth(root *TreeNode) int {
//    if root == nil{
//        return 0
//    }
//    left := maxDepth(root.Left)
//    right := maxDepth(root.Right)
//    return max(left, right)+1
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 从上向下 另一种方法
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var f func(p *TreeNode, n int)
	ans := 0
	f = func(p *TreeNode, n int) {
		if p == nil {
			return
		}
		n += 1
		ans = max(ans, n)
		f(p.Left, n)
		f(p.Right, n)
	}
	f(root, 0)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

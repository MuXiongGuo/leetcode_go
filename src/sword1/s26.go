package main

import "math"

// 树哈希
func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	m := map[*TreeNode]int{}
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left, depth+1), dfs(node.Right, depth+1)
		val := ((node.Val + left*13 + right*17) * depth) % math.MaxInt64
		m[node] = val
		return val
	}
	bVal := m[B]
	var test func(node *TreeNode) bool
	test = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if m[node] == bVal {
			return true
		}
		return test(node.Left) || test(node.Right)
	}
	return test(A)
}

// 不是树哈希 直接暴力就可以啦
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var helper func(a, b *TreeNode) bool
	helper = func(a, b *TreeNode) bool {
		if a == nil && b != nil {
			return false
		}
		if a != nil && b == nil {
			return true
		}
		if a == nil && b == nil {
			return true
		}

		if a.Val != b.Val {
			return false
		}
		return helper(a.Left, b.Left) && helper(a.Right, b.Right)
	}

	var dfs func(a, b *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		if a == nil {
			return false
		}
		return helper(a, b) || dfs(a.Left, b) || dfs(a.Right, b)
	}
	return dfs(A, B)
}

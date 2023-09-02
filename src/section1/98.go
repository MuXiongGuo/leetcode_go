package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 判断二叉搜索  核心还是递归, 设置最大最小 以及提前结束的变量 以及nil处理
// 没有bool也可以 失败的时候返回一个-inf inf 边界也可直接杀死(与nil反着)  后序
func isValidBST(root *TreeNode) bool {
	var helper func(p *TreeNode) (int, int, bool)
	helper = func(p *TreeNode) (int, int, bool) {
		if p == nil {
			return math.MaxInt64, math.MinInt64, true
		}
		minLVal, maxLVal, LFlag := helper(p.Left)
		minRVal, maxRVal, RFlag := helper(p.Right)
		if LFlag == false || RFlag == false || p.Val >= minRVal || p.Val <= maxLVal {
			return 0, 0, false
		}
		return min(minLVal, p.Val), max(maxRVal, p.Val), true // 加max min为了筛选掉nil
	}
	_, _, ans := helper(root)
	return ans
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

// 灵神将这个边界值放到了 参数里 func(p ,*TreeNode, small, big int) bool 这样也是不错的
// 思想是从上到下的 与 我那个翻着  对根节点没有要求  之后要求逐渐变高
// 前序思想
func isValidBST(root *TreeNode) bool {
	var helper func(p *TreeNode, left, right int) bool
	helper = func(p *TreeNode, left, right int) bool {
		if p == nil {
			return true
		}
		return p.Val < right && p.Val > left && helper(p.Left, left, p.Val) && helper(p.Right, p.Val, right)
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 中序 我们只要判断每个节点是严格递增即可
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	var helper func(p *TreeNode) bool
	helper = func(p *TreeNode) bool {
		if p == nil {
			return true
		}
		if !helper(p.Left) {
			return false
		}
		if p.Val <= pre {
			return false
		}
		pre = p.Val
		return helper(p.Right)
	}
	return helper(root)
}

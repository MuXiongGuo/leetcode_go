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

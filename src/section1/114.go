package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	l := []*TreeNode{}
	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		l = append(l, p)
		dfs(p.Left)
		dfs(p.Right)
	}
	dfs(root)
	for i := 0; i < len(l); i++ {
		l[i].Left = nil
		if i+1 < len(l) {
			l[i].Right = l[i+1]
		} else {
			l[i].Right = nil
		}
	}
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)
		ans = append(ans, []int{})
		for n > 0 {
			root = queue[0]
			queue = queue[1:]
			ans[len(ans)-1] = append(ans[len(ans)-1], root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			n--
		}
	}
	return ans
}

// 官方用dfs做的也很漂亮
// 但是逻辑上和bps是不同的 利用dep变量扩展深度,毕竟dfs擅长统计dep
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dep := 0
	var dfs func(head *TreeNode, dep int)
	dfs = func(head *TreeNode, dep int) {
		if head == nil {
			return
		}
		if len(res) == dep {
			res = append(res, []int{})
		}
		res[dep] = append(res[dep], head.Val)
		if head.Left != nil {
			dfs(head.Left, dep+1)
		}
		if head.Right != nil {
			dfs(head.Right, dep+1)
		}
	}
	dfs(root, dep)
	return res
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	var dfs func(head *TreeNode, dep int)
	dfs = func(root *TreeNode, dep int) {
		if root == nil {
			return
		}
		if dep == len(ans) {
			ans = append(ans, []int{})
		}
		ans[dep] = append(ans[dep], root.Val)
		dfs(root.Left, dep+1)
		dfs(root.Right, dep+1)
	}
	dfs(root, 0)
	return ans
}

// 官方的queue 用的太优雅了
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

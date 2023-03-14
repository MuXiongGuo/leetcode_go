package main

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	direction := "left"
	for len(queue) != 0 {
		n := len(queue)
		ans = append(ans, make([]int, n))
		if direction == "left" {
			direction = "right"
			for i := 0; i < n; i++ {
				root = queue[0]
				queue = queue[1:]
				ans[len(ans)-1][i] = root.Val
				if root.Left != nil {
					queue = append(queue, root.Left)
				}
				if root.Right != nil {
					queue = append(queue, root.Right)
				}
			}
		} else {
			direction = "left"
			for i := n - 1; i >= 0; i-- {
				root = queue[0]
				queue = queue[1:]
				ans[len(ans)-1][i] = root.Val
				if root.Left != nil {
					queue = append(queue, root.Left)
				}
				if root.Right != nil {
					queue = append(queue, root.Right)
				}
			}
		}
	}
	return ans
}

// 官方还是dfs 之后取那些偶数行逆序一下即可
// 灵活
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
	for i := 1; i < len(res); i += 2 {
		for j, n := 0, len(res[i]); j < n/2; j++ {
			res[i][j], res[i][n-1-j] = res[i][n-1-j], res[i][j]
		}
	}
	return res
}

// 官方的queue的用法 很牛 不用pop了都
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}

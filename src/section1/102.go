package main

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		n := len(q)
		tmp := []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
		ans = append(ans, tmp)
	}
	return ans
}

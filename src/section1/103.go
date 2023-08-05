package main

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	for i := 1; i < len(ans); i += 2 {
		reverse(ans[i])
	}
	return ans
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

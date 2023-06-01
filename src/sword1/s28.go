package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 别想太多 从上到下 只比较单个孩子的值与nil即可
func helper(p1 *TreeNode, p2 *TreeNode) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
		return false
	}
	if p1.Val == p2.Val {
		return helper(p1.Left, p2.Right) && helper(p1.Right, p2.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

// 官方 更优雅
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 迭代方法
func isSymmetric(root *TreeNode) bool {
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		u, v := q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)
		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
func main() {
	// [1,2,2,3,4,4,3]
	isSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}})
}

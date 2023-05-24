package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func BuildTree(s []int) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	// -1 为 nil节点   修改len(s) 可进行自定义 比如高度... 或已知总数之类的
	vTree := make([]*TreeNode, len(s))
	var root *TreeNode
	for i, el := range s {
		var node *TreeNode
		if el != -1 {
			node = &TreeNode{
				val:   el,
				left:  nil,
				right: nil,
			}
		}
		vTree[i] = node
	}
	root = vTree[0]
	for i := 0; 2*i+2 < len(s); i++ {
		if vTree[i] != nil {
			vTree[i].left = vTree[2*i+1]
			vTree[i].right = vTree[2*i+2]
		}
	}
	return root
}

func PreorderOutPutTree(p *TreeNode) {
	if p != nil {
		fmt.Print(p.val)
		fmt.Printf(" ")
		PreorderOutPutTree(p.left)
		PreorderOutPutTree(p.right)
	}
}

func InorderOutPutTree(p *TreeNode) {
	if p != nil {
		InorderOutPutTree(p.left)
		fmt.Print(p.val)
		fmt.Printf(" ")
		InorderOutPutTree(p.right)
	}
}

func PostorderOutPutTree(p *TreeNode) {
	if p != nil {
		PostorderOutPutTree(p.left)
		PostorderOutPutTree(p.right)
		fmt.Print(p.val)
		fmt.Printf(" ")
	}
}

func main() {
	p := BuildTree([]int{1, 2, 3, 4, 5, 6, -1}) // 要满一层才可以!!!
	PreorderOutPutTree(p)
	InorderOutPutTree(p)
	PostorderOutPutTree(p)
}

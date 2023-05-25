package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
				Val:   el,
				Left:  nil,
				Right: nil,
			}
		}
		vTree[i] = node
	}
	root = vTree[0]
	for i := 0; 2*i+2 < len(s); i++ {
		if vTree[i] != nil {
			vTree[i].Left = vTree[2*i+1]
			vTree[i].Right = vTree[2*i+2]
		}
	}
	return root
}

func PreorderOutPutTree(p *TreeNode) {
	if p != nil {
		fmt.Print(p.Val)
		fmt.Printf(" ")
		PreorderOutPutTree(p.Left)
		PreorderOutPutTree(p.Right)
	}
}

func InorderOutPutTree(p *TreeNode) {
	if p != nil {
		InorderOutPutTree(p.Left)
		fmt.Print(p.Val)
		fmt.Printf(" ")
		InorderOutPutTree(p.Right)
	}
}

func PostorderOutPutTree(p *TreeNode) {
	if p != nil {
		PostorderOutPutTree(p.Left)
		PostorderOutPutTree(p.Right)
		fmt.Print(p.Val)
		fmt.Printf(" ")
	}
}

func main() {
	p := BuildTree([]int{1, 2, 3, 4, 5, 6, -1}) // 要满一层才可以!!!
	PreorderOutPutTree(p)
	//InorderOutPutTree(p)
	//PostorderOutPutTree(p)
}

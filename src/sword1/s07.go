package main

import (
	"fmt"
)

//	type TreeNode struct {
//		Val   int
//		Left  *TreeNode
//		Right *TreeNode
//	}
//
// 小心 指针的引用 已经不能用了 指针的指针才可以 但是可以巧妙地避免

func buildTree(preorder []int, inorder []int) *TreeNode {
	find := func(s []int, x int) (ans int) {
		for i, v := range s {
			if x == v {
				return i
			}
		}
		return
	}
	var helper func(pre, in []int, p *TreeNode)
	helper = func(pre, in []int, p *TreeNode) {
		// 终止条件
		if len(pre) == 0 {
			return
		}
		p.Val = pre[0]
		index := find(inorder, pre[0])
		// 开辟新节点
		if len(pre[1:index+1]) != 0 {
			p.Left = &TreeNode{}
		}
		if len(pre[index+1:]) != 0 {
			p.Right = &TreeNode{}
		}
		helper(pre[1:index+1], in[:index], p.Left)   // 单个元素s[1:] 不报错
		helper(pre[index+1:], in[index+1:], p.Right) // 长度3  s[3:] 也不报错 很有趣的性质
	}
	if len(preorder) == 0 {
		return nil
	}
	p := &TreeNode{}
	helper(preorder, inorder, p)
	return p
}

// 官方!!! 思路是一样的 但是完美利用了返回值!!!
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func main() {
	// 一个空的 cap 为 5 的 slice
	queue := make([]int, 0, 5) // _ _ _ _ _
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 0, cap(queue): 5, queue: []

	// 入队一个元素
	queue = append(queue, 1) // [1] _ _ _ _
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 1, cap(queue): 5, queue: [1]

	// 再入队一个元素
	queue = append(queue, 2) // [1 2] _ _ _
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 2, cap(queue): 5, queue: [1 2]

	// 出队一个元素
	queue = queue[1:] // 1 [2] _ _ _
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 1, cap(queue): 4, queue: [2]

	// 再入队三个元素
	queue = append(queue, 3, 4, 5) // 1 [2 3 4 5]
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 4, cap(queue): 4, queue: [2 3 4 5]

	// 出队二个元素
	queue = queue[2:] // 1 2 3 [4 5]
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 2, cap(queue): 2, queue: [4 5]

	// 再入队一个元素，触发 slise 的扩容，根据扩容策略，此时 slice cap 为 2，翻倍为 4
	queue = append(queue, 6) //  // [4 5 6] _
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
	// 输出 len(queue): 3, cap(queue): 4, queue: [4 5 6]

	queue = queue[:len(queue)-1] //  // [4 5 6] _
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)

	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	var s3 []int
	fmt.Println(s2)
	fmt.Println(s3)
	s4 := s1[3:]
	//p := &TreeNode{}
	//fmt.Println(p)
	fmt.Println(s4)
}

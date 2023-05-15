package main

import "fmt"

//	func buildTree(preorder []int, inorder []int) *TreeNode {
//		var helper func(start, end int, p *TreeNode)
//		helper = func(start, end int, p *TreeNode) {
//			if start > end {
//				return
//			}
//			p.Val = preorder[start]
//
//		}
//	}
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

}

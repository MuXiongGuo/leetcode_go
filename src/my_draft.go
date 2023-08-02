package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

func main() {
	//s1 := []int{1, 2, 3, 4}
	//s2 := []int{9, 4}
	//s1, s2 = s2, s1
	//fmt.Println(s1)
	//fmt.Println(s2)

	// 类型别名的用法
	//type mm []int
	//s := []int{1, 23}
	//var m mm
	//m = s
	//fmt.Println(m)

	//heap的使用
	h := make(minHeap, 0)
	heap.Init(&h)

	heap.Push(&h, 2)
	heap.Push(&h, 7)
	heap.Push(&h, 5)
	heap.Push(&h, 5)
	heap.Push(&h, 4)
	heap.Push(&h, -6)

	fmt.Println(h[0])
	//
	//for len(h) != 0 {
	//	fmt.Println(heap.Pop(&h))
	//}
	//x := iner{1, 2, 3}
	////x.re()
	////fmt.Println(*x)
	//var xx myer = x
	////var xx myer = x
	//x[1] = 100
	//var xxx myer = xx
	//xxx.fmt.Println(xx)
	//x := iner{1, 2, 3}
	//var xx myer = &x
	//xx.re()
	//fmt.Println(xx)
	//fmt.Println(x)
	// (1) 指针类型 或 值类型 无论值还是指针去调用效果一样  智能选择!
	// (2) 接口是某个的变量的& 那就是指针类型实现了所需要的接口  且不做copy 直接引用
	// (3) 若是非& 则是一份copy 无法修改!

	// 获得切片的元素类型
	v1 := reflect.ValueOf(minHeap{2, 3})
	v2 := v1.Type().Elem()
	fmt.Println(v2)

}

type iner [3]int

func (i iner) re() {
	i[0] = 10
}

type myer interface {
	re()
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

// 这里实现了小根堆，如果想要实现大根堆可以改为 h[i] > h[j]
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	(h)[i], (h)[j] = (h)[j], (h)[i]
}

func (h *minHeap) Push(x interface{}) {
	reflect.TypeOf(h)
	//*h = append(*h, x.(int))
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

//https://blog.csdn.net/qq_42696069/article/details/129397728
//https://blog.csdn.net/qq_42696069/article/details/129397728

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	dummyNode := &ListNode{Next: head}
	n := 0
	for p := dummyNode.Next; p != nil; p = p.Next {
		n++
	}
	start := dummyNode
	for n >= k {
		n -= k
		q0, q1 := start.Next, start.Next.Next
		var q2 *ListNode
		for i := 1; i < k; i++ {
			q2 = q1.Next
			q1.Next = q0
			q0 = q1
			q1 = q2
		}
		// 此时 q0为尾部   q1=q2 为q0下一个节点
		q2 = start.Next
		start.Next.Next = q1
		start.Next = q0
		start = q2
	}
	return dummyNode.Next

}
func main() {
	reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	iner.re()
}

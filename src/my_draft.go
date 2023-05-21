package main

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

	// heap的使用
	//h := make(minHeap, 0)
	//heap.Init(&h)
	//
	//heap.Push(&h, 2)
	//heap.Push(&h, 7)
	//heap.Push(&h, 5)
	//heap.Push(&h, 5)
	//heap.Push(&h, 4)
	//heap.Push(&h, 6)
	//
	//for len(h) != 0 {
	//	fmt.Println(heap.Pop(&h))
	//}
	//x := iner{1, 2, 3}
	////x.re()
	////fmt.Println(*x)
	//var xx myer = &x
	//var xx myer = x
	//x[1] = 100
	//fmt.Println(xx)

	// (1) 指针类型 或 值类型 无论值还是指针去调用效果一样  智能选择!
	// (2) 接口是某个的变量的& 那就是指针类型实现了所需要的接口  且不做copy 直接引用
	// (3) 若是非& 则是一份copy 无法修改!

}

type iner [3]int

func (i iner) re() {
	i[0] = 10
}

type myer interface {
	re()
}

//
//type minHeap []int
//
//func (h minHeap) Len() int {
//	return len(h)
//}
//
//// 这里实现了小根堆，如果想要实现大根堆可以改为 h[i] > h[j]
//func (h minHeap) Less(i, j int) bool {
//	return h[i] < h[j]
//}
//
//func (h *minHeap) Swap(i, j int) {
//	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
//}
//
//func (h *minHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//
//func (h *minHeap) Pop() interface{} {
//	res := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return res
//}

//https://blog.csdn.net/qq_42696069/article/details/129397728
//https://blog.csdn.net/qq_42696069/article/details/129397728

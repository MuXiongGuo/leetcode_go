package main

import "fmt"

// 切片 多维切片共享
func main() {
	var ans [][]int
	var ans2 [][]int
	a := []int{2, 3, 4}
	b := []int{12, 55, 44}
	ans = append(ans, a)
	ans[0][1] = 100
	fmt.Println(a)
	ans[0] = b
	ans[0][2] = 200
	fmt.Println(b)
	fmt.Println(ans)

	//ans2 = append(ans2, ans[0])
	//ans2[0][0] = 999
	//fmt.Println(ans)
	ans2 = append(ans2, []int{ans[0][0], ans[0][1]})
	ans2[0][0] = 999
	fmt.Println(ans)
}

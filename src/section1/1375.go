package main

import "fmt"

// 二进制数学而已!  状态压缩
// 一个坑! 数据太长会溢出 想办法解决!
func numTimesAllBlue(flips []int) int {
	n, ret, s, obj := len(flips), 0, 0, 0
	for i, x := range flips {
		obj += 1 << (n - i - 1)
		s += 1 << (n - x)
		if obj == s {
			ret++
		}
	}
	return ret
}

// 官方 之前想复杂了，本质是转换思路，第i次的最大值为i即可 只记录下最大值就行了
func numTimesAllBlue(flips []int) int {
	res, s := 0, 0
	for i, v := range flips {
		s = max(s, v)
		if s == i+1 {
			res++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5}))

}

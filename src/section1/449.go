package main

import (
	"math"
	"strconv"
	"strings"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.

func (this *Codec) serialize(root *TreeNode) string {
	var ans strings.Builder
	var helper func(p *TreeNode)
	helper = func(p *TreeNode) {
		if p == nil {
			return
		}
		ans.WriteString(strconv.Itoa(p.Val))
		ans.WriteString("#")
		helper(p.Left)
		helper(p.Right)
	}
	helper(root)
	var s string
	if len(ans.String()) != 0 {
		s = ans.String()[:len(ans.String())-1]
	}
	return s
}

// Deserializes your encoded data to tree.

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	dataSlice := strings.Split(data, "#")
	dataIntSlice := make([]int, len(dataSlice))
	for id, el := range dataSlice {
		dataIntSlice[id], _ = strconv.Atoi(el)
	}
	var helper func(s []int) *TreeNode
	helper = func(s []int) *TreeNode {
		if len(s) == 0 {
			return nil
		}
		p := &TreeNode{Val: s[0]}
		idx := 1
		for ; idx < len(s); idx++ {
			if s[idx] > s[0] {
				break
			}
		}
		p.Left = helper(s[1:idx])
		p.Right = helper(s[idx:])
		return p
	}
	return helper(dataIntSlice)
}

/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor()
* deser := Constructor()
* tree := ser.serialize(root)
* ans := deser.deserialize(tree)
* return ans
 */

//func main() {
//	a := strings.Builder{}
//	a.WriteString("dadas")
//	fmt.Println(a.String())
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [2, 1, 3]
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	ser := Constructor()
	deser := Constructor()
	tree := ser.serialize(root)
	ans := deser.deserialize(tree)
	println(ans.Val)

}

// 官方 利用了[]string 以及join函数 比较聪明 并且利用了判断边界来结束 而不是用切割
type Codec struct{}

func Constructor() (_ Codec) { return }

func (Codec) serialize(root *TreeNode) string {
	arr := []string{}
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		arr = append(arr, strconv.Itoa(node.Val))
	}
	postOrder(root)
	return strings.Join(arr, " ")
}

func (Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var construct func(int, int) *TreeNode
	construct = func(lower, upper int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[:len(arr)-1]
		return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

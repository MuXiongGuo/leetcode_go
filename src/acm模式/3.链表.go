package main

import "fmt"

type List struct {
	Val  int
	Next *List
}

func CreateList(s []int) *List {
	if len(s) == 0 {
		return nil
	}
	ret := &List{
		Val:  s[0],
		Next: nil,
	}
	p := ret
	for i := 1; i < len(s); i++ {
		p.Next = &List{Val: s[i], Next: nil}
		p = p.Next
	}
	return ret
}

func OutPutList(p *List) {
	for p != nil {
		fmt.Print(p.Val)
		fmt.Printf(" ")
		p = p.Next
	}
}

func main() {
	p := CreateList([]int{2, 3, 4, 9})
	OutPutList(p)
}

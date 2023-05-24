package main

import "fmt"

type List struct {
	val  int
	next *List
}

func CreateList(s []int) *List {
	if len(s) == 0 {
		return nil
	}
	ret := &List{
		val:  s[0],
		next: nil,
	}
	p := ret
	for i := 1; i < len(s); i++ {
		p.next = &List{val: s[i], next: nil}
		p = p.next
	}
	return ret
}

func OutPutList(p *List) {
	for p != nil {
		fmt.Print(p.val)
		fmt.Printf(" ")
		p = p.next
	}
}

func main() {
	p := CreateList([]int{2, 3, 4, 9})
	OutPutList(p)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func InputRowRetInt() []int {
	var intS []int
	sc.Scan()
	strS := strings.Split(sc.Text(), " ")
	for _, el := range strS {
		val, _ := strconv.Atoi(el)
		intS = append(intS, val)
	}
	return intS
}

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
	// 读两行 读多行用for 即可
	s1 := InputRowRetInt()
	s2 := InputRowRetInt()

	p1 := CreateList(s1)
	p2 := CreateList(s2)
	var p3 = &List{}
	head := p3
	for p1 != nil && p2 != nil {
		p3.Next = &List{}
		if p1.Val < p2.Val {
			p3.Next.Val = p1.Val
			p1 = p1.Next
		} else {
			p3.Next.Val = p2.Val
			p2 = p2.Next
		}
		p3 = p3.Next
	}
	for p1 != nil {
		p3.Next = &List{}
		p3.Next.Val = p1.Val
		p1 = p1.Next
		p3 = p3.Next
	}
	for p2 != nil {
		p3.Next = &List{}
		p3.Next.Val = p2.Val
		p2 = p2.Next
		p3 = p3.Next
	}
	OutPutList(head.Next)
}

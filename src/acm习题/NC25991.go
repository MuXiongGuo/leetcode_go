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
	// 读两行 读多行用for 即可
	s1 := InputRowRetInt()
	s2 := InputRowRetInt()

	p1 := CreateList(s1)
	p2 := CreateList(s2)
	var p3 = &List{}
	head := p3
	for p1 != nil && p2 != nil {
		p3.next = &List{}
		if p1.val < p2.val {
			p3.next.val = p1.val
			p1 = p1.next
		} else {
			p3.next.val = p2.val
			p2 = p2.next
		}
		p3 = p3.next
	}
	for p1 != nil {
		p3.next = &List{}
		p3.next.val = p1.val
		p1 = p1.next
		p3 = p3.next
	}
	for p2 != nil {
		p3.next = &List{}
		p3.next.val = p2.val
		p2 = p2.next
		p3 = p3.next
	}
	OutPutList(head.next)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	// 读两行 读多行用for 即可
	sc := bufio.NewScanner(os.Stdin)
	list1 := []int{}
	sc.Scan()
	strs1 := strings.Split(sc.Text(), " ")
	for _, el := range strs1 {
		val, _ := strconv.Atoi(el)
		list1 = append(list1, val)
	}
	sc.Scan()
	strs2 := strings.Split(sc.Text(), " ")
	for _, el := range strs2 {
		val, _ := strconv.Atoi(el)
		list1 = append(list1, val)
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	// 输出格式
	for _, el := range list1 {
		s := strconv.Itoa(el)
		fmt.Printf(s)
		fmt.Printf(" ")
	}
}

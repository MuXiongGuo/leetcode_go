package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
	for _, el := range list1 {
		s := strconv.Itoa(el)
		fmt.Printf(s)
		fmt.Printf(" ")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func OutPutIntSlice(s []int) {
	// 空格间隔 有结尾空格
	//for _, el := range s {
	//	val := strconv.Itoa(el)
	//	fmt.Printf(val)
	//	fmt.Printf(" ")
	//}

	// 无结尾空格
	for i := 0; i < len(s); i++ {
		val := strconv.Itoa(s[i])
		fmt.Printf(val)
		if i != len(s)-1 {
			fmt.Printf(" ")
		}
	}
}

func main() {
	// 读两行 读多行用for 即可
	s1 := InputRowRetInt()
	s2 := InputRowRetInt()
	s1 = append(s1, s2...)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	// 输出格式
	OutPutIntSlice(s1)
}

package main

import "fmt"

// 6. 学习匿名字段
//type data struct {
//	*int
//	string
//}

// 7. 学习结构体
type base1 struct {
	content string
}

type base2 struct {
	b1  base1
	con string
}

func main() {
	// 1. range copy问题
	//a := []int{1, 2, 3, 4, 5}
	//for _, v := range a {
	//	v += 10
	//}
	//fmt.Println(a)

	// 2. 匿名函数外部变量影响
	//a := 10
	//f := func() {
	//	a += 10
	//}
	//f()
	//println(a)

	// 3. 切片追加问题
	//s := make([]int, 0, 100)
	//s[1:5][0] = 99
	//fmt.Println(s[:10])
	//s1 := s[1:5:10]
	//s1 = append(s1, 1)
	//fmt.Println(s1)
	//fmt.Println(s[:10])

	// 4. 字典 Initialization
	//m := map[string]int{
	//	"a": 3,
	//}
	//m["a"] += 1
	//fmt.Println(m)

	// 5. range copy问题
	//count := 0
	//a := []int{1, 2, 3, 4, 5}
	//for i, _ := range a {
	//	if i == 0 {
	//		a = append(a, 6)
	//		a = append(a, 7)
	//		a = append(a, 8)
	//	}
	//	count += 1
	//}
	//println(count)

	// 6. 学习匿名字段
	//x := 100
	//d := data{
	//	int:    &x,
	//	string: "ddddcccc",
	//}
	//d.string = "aaa"
	//fmt.Println(d)

	// 7. 学习结构体
	b := base2{
		con: "dddaaaa",
		b1: base1{
			content: "p",
		},
	}
	b.b1.content = "daio"
	fmt.Println(b)
}

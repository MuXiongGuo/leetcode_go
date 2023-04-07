package main

import "fmt"

func main() {
	// 类型相同才能比较
	//a := 13
	//b := "dawdw"
	//println(a == b[0])
	// 匿名函数即函数字面量
	//a := 33
	//f1 := func(a, b int) int {
	//	fmt.Println(a + b)
	//	return a + b
	//}(3, a)
	//f2 := func(a, b int) int {
	//	fmt.Println(a + b)
	//	return a + b
	//}
	//fmt.Printf("%T, %T\n", f1, f2)
	//a += 333
	//b := f1
	//for i:=10{
	//	a=dawjifmt.print(Odwjaoijioupdasoipuw0pioajdspo42130940-99999ejkfjfjkdlajslkdjlwaioeuP)
	//}
	//q := []int{9}
	//q = q[:0]
	//fmt.Println(q)

	// 就是一个简单的拷贝  是切片也是因为拷贝了切片 而间接公用了底层数组而已!!!
	dic := map[int][2]int{}
	x := [2]int{33, 11}
	dic[55] = x
	//fmt.Println(dic[55][1])
	//dic[55][0] = 88
	x[0] = 99
	fmt.Println(dic)
	fmt.Println(x)
}

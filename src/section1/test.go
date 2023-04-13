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

	// 就是一个简单的拷贝  是切片也是因为拷贝了切片 而间接公用了底层数组而已!!!  字典的修改
	dic := map[int][2]int{}
	x := [2]int{33, 11}
	dic[55] = x
	//fmt.Println(dic[55][1])
	//dic[55][0] = 88
	x[0] = 99
	fmt.Println(dic)
	fmt.Println(x)
	dic2 := map[int][]int{}
	x2 := []int{22, 44}
	dic2[18] = x2
	dic2[18][0] = 88
	x2[1] = 444
	fmt.Println(dic2)

	//
	c := make(chan int, 2)
	c <- 5
	select {
	case <-c:
		fmt.Println("end")
	}

	// string
	//	sNew := `ni hao wo shi
	//hhahaha\ndddddd`
	//	sNew2 := "ni hao wo shinhhadaha\n" +
	//		"xin de yi hang"
	//	requestBuf := "GET /go HTTP/1.1\r\n" +
	//		"Accept: image/gif, image/jpeg,image/pjpeg, application/x-ms-application, application/xaml+xml,application/x-ms-xbap, */*\r\n" +
	//		"Accept-Language:zh-Hans-CN,zh-Hans;q=0.8,en-US;q=0.5,en;q=0.3\r\n" +
	//		"User-Agent:Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64;Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR3.0.30729; .NET CLR 3.5.30729)\r\n" +
	//		"Accept-Encoding: gzip,deflate\r\n" +
	//		"Host: 127.0.0.1:8000\r\n" +
	//		"Connection: Keep-Alive\r\n\r\n"
	//	fmt.Println(sNew)
	//	fmt.Println(sNew2)
	//	fmt.Println(requestBuf)

	n1, m1 := test()
	fmt.Println(n1)
	fmt.Println(m1)
}

func test() (a, b int) {
	if true {
		a := 3
		b := 4
		fmt.Println(a + b)
	}
	return
}

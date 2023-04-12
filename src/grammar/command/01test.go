package main

// 禁止内联输出优化信息
// go build -gcflags "-l -m" 01test.go
// 反汇编确认
// go tool objdump -s "main.main" 01test.exe

// 默认允许内联优化

func test() *int {
	a := 0x100
	return &a
}

func main() {
	var x *int = test()
	println(x, *x)
}

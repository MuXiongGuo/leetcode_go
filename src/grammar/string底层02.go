package main

func main() {
	// 两种情况下编译器会优化string 和 byte转换避免额外的底层复制
	// 1. 将[]byte转为string key, 去查询map[string]时
	// 2. 将string转为[]byte, 进行for range迭代时, 直接取字节赋值给局部变量

	m := map[string]int{
		"hello": 100,
	}
	// gdb调试
	key := []byte("hello")
	x, ok := m[string(key)] // 无底层copy的转换
	println(x, ok)
}

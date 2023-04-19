package main

import (
	"fmt"
	"testing"
	"time"
)

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(3, 4))
	// Output:
	// 3
	// 7
}

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("Add(1, 2) != 3")
	}
}

func TestA(t *testing.T) {
	t.Parallel() // 两个都加的话就能并行测试 大大加速!
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	t.Parallel() // 两个都加的话就能并行测试 大大加速!
	time.Sleep(time.Second * 2)
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs() // 输出内存信息
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func heap() []byte {
	return make([]byte, 1024*1024)
}

func BenchmarkHeap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap()
	}
}

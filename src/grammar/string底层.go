package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 切片时使用的仍是一个底层
	s := "abcdefg"
	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[1:2]
	fmt.Printf("%#v\n", &s)
	fmt.Printf("%#v\n", &s3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s2)))

	// byte rune转化时  会造成底层的copy
	b := []byte(s)
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
	sB := string(b)
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&sB)))

	// 黑魔法效果
	bNew := String2Bytes(s)
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&bNew)))
	fmt.Printf("%#v\n", &bNew)
	fmt.Printf("%#v\n", &s)
	// 其他知识
	//sNew := []byte("abcdefg")
	//sss := sNew...
}

// 黑魔法

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

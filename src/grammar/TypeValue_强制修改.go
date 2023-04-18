package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	code int
}

func main() {
	p := new(User) // 要用指针 因为接口变量会复制对象, 且是unaddressable, 想修改目标就要用指针
	v := reflect.ValueOf(p).Elem()

	name := v.FieldByName("Name")
	code := v.FieldByName("code")

	fmt.Println("name: canAddr = %v, canSet = %v\n", name.CanAddr(), name.CanSet())
	fmt.Println("code: canAddr = %v, canSet = %v\n", code.CanAddr(), code.CanSet())

	if name.CanSet() {
		name.SetString("Tom")
	}
	if code.CanAddr() {
		//code.SetInt(100)  不可修改 可以想办法强制修改
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}
	fmt.Printf("p = %+v\n", *p)
}

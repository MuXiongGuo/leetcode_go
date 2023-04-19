package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	// 1
	println(a == nil) // true
	println(b == nil) // false
	//println(reflect.ValueOf(a).IsNil()) 不能这样写!!!
	println(reflect.ValueOf(b).IsNil()) // true

	// 2
	//cInt := 99
	var c interface{} = (*int)(nil)
	uc := unsafe.Pointer(&c)
	fmt.Println(uc, *(*uintptr)(uc)) // 0xc0000a4008 0
	iface := (*[2]uintptr)(unsafe.Pointer(&c))
	fmt.Println(iface, iface[1] == 0) // [0 0]
}

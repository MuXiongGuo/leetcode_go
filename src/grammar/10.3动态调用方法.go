package main

import (
	"fmt"
	"reflect"
)

// 反射是实现元编程的重要手段
// 动态调用方法 传递参数

type X struct{}

func (*X) Test(x, y int) (int, error) {
	return x + y, nil
}

func (*X) Test2(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

func main() {
	var x X
	v := reflect.ValueOf(&x)
	m := v.MethodByName("Test")

	in := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	out := m.Call(in)
	fmt.Println(out)
	for _, v := range out {
		fmt.Println(v)
	}

	// 可变参数推荐用CallSlice 只要一个切片就能包含args所有额外的
	m2 := v.MethodByName("Test2")
	out1 := m2.Call([]reflect.Value{
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf("age"),
		reflect.ValueOf(18),
	})

	out2 := m2.CallSlice([]reflect.Value{
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf([]interface{}{"age", 18}),
	})

	fmt.Println(out1)
	fmt.Println(out2)
}

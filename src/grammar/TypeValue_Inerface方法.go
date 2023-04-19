package main

import (
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		Name string
		Age  int
	}

	u := user{"Tom", 20}

	v := reflect.ValueOf(&u)

	p := v.Interface()
	fmt.Println(p)
	empty := interface{}(&user{
		Name: "aa",
		Age:  10,
	})
	empty.(*user).Name = "bb" // 接口回退

	pNew := v.Interface().(*user)
	fmt.Println(pNew)
	fmt.Println(empty)
}

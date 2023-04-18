package main

import "reflect"

// 获取结构体指针的基类型 才能 遍历字段

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 输出匿名字段的信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			for j := 0; j < f.Type.NumField(); j++ {
				af := f.Type.Field(j)
				println(af.Name, af.Type, af.Offset)
			}
		}
	}
}

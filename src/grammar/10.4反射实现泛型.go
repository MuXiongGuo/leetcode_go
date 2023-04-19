package main

import (
	"reflect"
	"strings"
)

// Add 通用加法
func Add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value

	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, len(args))
		for i, a := range args {
			ss[i] = a.String()
		}
		ret = reflect.ValueOf(strings.Join(ss, ""))
	}
	results = append(results, ret)
	return
}

// MakeAdd 将函数指针参数指向通用算法函数
func MakeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), Add) // 关键
	fn.Set(v)                             // 指向通用算法函数
}

func main() {
	var intAdd func(int, int, int) int
	MakeAdd(&intAdd)
	println(intAdd(1, 2, 9))

	var strAdd func(string, string) string
	MakeAdd(&strAdd)
	println(strAdd("hello", "world"))
}

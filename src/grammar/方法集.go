package main

// 这个为啥不能用!!!  很奇怪
import (
	"fmt"
	"reflect"
)

type Sbig struct{}
type Tsmall struct {
	age  int
	work string
}

//	func (s Sbig) sVal() {
//		fmt.Println("hello")
//	}
//
//	func (s *Sbig) sPtr() {
//		fmt.Println("hello")
//	}
func (t Tsmall) TVal() {
	fmt.Println("hello")
}
func (t *Tsmall) TPtr() {
	fmt.Println("hello")
}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t Tsmall = Tsmall{age: 1, work: "2"}
	methodSet(t)
	println("------")
	methodSet(&t)
}

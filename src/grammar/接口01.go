package main

import "fmt"

type T struct {
	a int
}

func (t *T) add() {
	t.a++
}

type ter interface {
	add()
}

func main() {
	//t := T{}
	//var el ter = &t
	//el.add()
	//fmt.Println(t.a)
	//fmt.Println(el.(*T).a)
	t := T{}
	var el interface{} = t
	t.a++
	fmt.Println(el.(T).a)
	fmt.Println(t.a)
}

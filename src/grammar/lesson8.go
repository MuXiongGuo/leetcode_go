package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Teacher struct {
	Person
	id int
}

func (t *Teacher) ChangePointer() {
	t.age = 77
}
func (t *Teacher) ChangeValue() {
	t.age = 88
}
func (t *Person) ChangePersonPointer() {
	t.age = 111
}

type Hummaner interface {
	SayHello()
}

func main() {
	p := Person{
		name: "Tom",
		age:  10,
	}
	t := Teacher{
		Person: p,
		id:     0,
	}
	pT := &Teacher{
		Person: Person{"jack", 22},
		id:     100,
	}
	t.age = 99
	// 指针和对象都是可以使用这个函数的
	//t.ChangeValue()
	//pT.ChangeValue()
	pT.ChangePointer()
	t.ChangePointer()
	t.ChangePersonPointer()
	fmt.Println(t)
	hSlice := make([]Hummaner, 10)
	fmt.Println(hSlice)
	//fmt.Println(p)
}

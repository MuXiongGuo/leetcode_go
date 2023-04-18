package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	X string
	Y int
}

func (f *Foo) Do() {
	//fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)
	fmt.Println("Do()\r\n")
}

func (f Foo) Doa(str string) {
	//fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)
	fmt.Println("Doa()\r\n", str)
}

func (f Foo) Dob() {
	//fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)
	fmt.Println("Dob()\r\n")
}

func (f *Foo) Doc() {
	//fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)
	fmt.Println("Doc()\r\n")
}

func main() {
	var f Foo
	typ := reflect.TypeOf(f)
	typa := reflect.TypeOf(&f)

	fmt.Println(typ.NumMethod(), typa.NumMethod()) //2  4
	m := typ.Method(0)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)
	fmt.Println(m.Index)

	m = typ.Method(1)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)
	fmt.Println(m.Index)

	fmt.Println("___________________________________________________________________\r\n")

	m = typa.Method(0)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)
	fmt.Println(m.Index)

	m = typa.Method(1)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)
	fmt.Println(m.Index)

	m = typa.Method(2)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)
	fmt.Println(m.Index)

	m = typa.Method(3)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)
	fmt.Println(m.Index)
}

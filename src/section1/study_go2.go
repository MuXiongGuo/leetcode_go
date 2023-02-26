package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	var p Point         //正常
	fmt.Println(p)      //输出 {0 0}
	var nilPtr *Point   //报错
	fmt.Println(nilPtr) //<nil>
	var newPtr *Point = new(Point)
	fmt.Println(newPtr) //&{0 0}

	//调用Abs()
	fmt.Println(p.Abs()) //正常
	// fmt.Println(nilPtr.Abs()) //报错
	// fmt.Println(newPtr.Abs()) //正常

}

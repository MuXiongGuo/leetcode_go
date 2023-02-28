package main

type data int

func (d data) p() {
	println(d)
}
func main() {
	var x int
	x := 22
	var name data
	name = 99
	name.p()
	//	x.p()
	x = int(name)
}

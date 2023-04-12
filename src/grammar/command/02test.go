package main

func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	f := test(132)
	f()
}

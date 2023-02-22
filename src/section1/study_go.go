package main

func main() {
	m := map[string]int{
		"abc": 13,
	}
	x, ok := m["abc"]
	println(x, ok)
}

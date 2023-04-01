package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func main() {
	t1 := IT{
		Company:  "TC",
		Subjects: []string{"python", "c++"},
		IsOk:     false,
		Price:    1400,
	}
	b, _ := json.Marshal(t1) // b此时是一个unit8的数组
	fmt.Println(string(b))
}

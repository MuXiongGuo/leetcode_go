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

type IT1 struct {
	Company  string   `json:"-"`
	Subjects []string `json:"subjects"`
	isOk     bool
	Price    float64 `json:"price,omitempty"`
}

type UIT struct {
	Company string
	Age     int // age不在json里面也能去偷Company, 大小写问题没太去考虑
}

func main() {
	// 编码部分
	t1 := IT{
		Company:  "TC",
		Subjects: []string{"python", "c++"},
		IsOk:     false,
		Price:    1400,
	}
	t2 := IT1{
		Company:  "hello",
		Subjects: []string{"python", "c++"},
		isOk:     false,
		Price:    55,
	}
	b, _ := json.Marshal(t1) // b此时是一个unit8的数组
	c, _ := json.Marshal(t2)
	fmt.Println(string(b))
	fmt.Println(string(c))
	// 解码部分
	var ans UIT
	err := json.Unmarshal(b, &ans)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)
}

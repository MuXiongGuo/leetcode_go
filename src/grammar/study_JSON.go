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
	// JSON解析到空接口中, 再配合断言将其读出
	bJson := []byte(`{
	"company":"itcast",
	"subjects":["Go", "C++", "Java"],
	"isok":true,
	"price":66.66
	}`)
	var tJson interface{}
	errJson := json.Unmarshal(bJson, &tJson)
	if errJson != nil {
		fmt.Println(errJson)
	}
	fmt.Println(tJson)
	// 使用断言
	mJson := tJson.(map[string]interface{})
	fmt.Println(mJson)
	for _, v := range mJson {
		// switch里面专用的类型判断, vv 就是v
		// 可以直接switch v.type
		switch vv := v.(type) {
		case string:
			fmt.Println(v)
		case []interface{}:
			fmt.Println(vv)
		}
	}
}

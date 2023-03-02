package main

import (
	"math"
	"strings"
)

func printBin(num float64) string {
	if num == 0.0 {
		return "0"
	}
	s := "0."
	n := int(math.Ceil(math.Log2(1 / 1e-6)))
	for i := 0; i < n; i++ {
		num *= 2
		if num == 1.0 {
			s += "1"
			return s
		} else if num > 1.0 {
			s += "1"
			num -= 1
		} else {
			s += "0"
		}
	}
	return "ERROR"
}

// 官方的 利用byte 截取前面的位数 省去了判断
func printBin(num float64) string {
	sb := &strings.Builder{}
	sb.WriteString("0.")
	for sb.Len() <= 32 && num != 0 {
		num *= 2
		digit := byte(num)
		sb.WriteByte('0' + digit)
		num -= float64(digit)
	}
	if sb.Len() <= 32 {
		return sb.String()
	}
	return "ERROR"
}

func main() {
	println(printBin(0.625))
	c := 0.254
	d := byte(c)
	println(d)
	println('0' + 55)
}

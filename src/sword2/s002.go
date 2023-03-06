package main

import "unsafe"

// answer = x ^ y  异或
// carry = (x & y) << 1
// x = answer  y = carry
// x y为ab的整数int格式
// python3 位运算模拟加法 看官方题解
//class Solution:
//    def addBinary(self, a, b) -> str:
//        x, y = int(a, 2), int(b, 2)
//        while y:
//            answer = x ^ y
//            carry = (x & y) << 1
//            x, y = answer, carry
//        return bin(x)[2:]
//

func addBinary(a string, b string) string {
	var ans []byte
	n, m := len(a), len(b)
	c := 0
	for i, j := n-1, m-1; i >= 0 || j >= 0; {
		val := c
		if i < 0 {
			val += int(b[j]) - '0'
		} else if j < 0 {
			val += int(a[i]) - '0'
		} else {
			val += int(a[i]) - '0' + int(b[j]) - '0'
		}
		switch val {
		case 0:
			c = 0
			ans = append(ans, '0')
		case 1:
			c = 0
			ans = append(ans, '1')
		case 2:
			c = 1
			ans = append(ans, '0')
		case 3:
			c = 1
			ans = append(ans, '1')
		}
		i--
		j--
	}
	if c == 1 {
		ans = append(ans, '1')
	}
	// reverse
	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	//return string(ans)
	return *(*string)(unsafe.Pointer(&ans)) // 无底层copy的转换
}

func main() {
	addBinary("11", "10")
}

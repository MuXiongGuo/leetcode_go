package main

import "math"

// 只能存32位 要注意溢出  以及 都转成负数
// 左移右移
func divide(a int, b int) int {
	if a == 0 {
		return 0
	}
	if a == b {
		return 1
	}
	if a == -b {
		return -1
	}

	var symbol int
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		symbol = -1
	} else {
		symbol = 1
	}

	a, b = int(math.Abs(float64(a))), int(math.Abs(float64(b)))
	if a < b {
		return symbol
	}

	ans := 0
	for {
		x, y := 1, b
		for y < a {
			x, y = x<<1, y<<1
		}
		if y == a {
			ans += x
			break
		} else {
			// 不可能再继续乘法了
			ans += x >> 1
			a -= y >> 1
			if a < b {
				break
			}
		}
	}
	ans *= symbol
	if ans >= 1<<31 || ans < -1<<31 {
		return (1 << 31) - 1
	}
	return ans
}

// 官方1    Z×Y≥X>(Z+1)×Y  聪明!
func main() {
	divide(7, -3)
}

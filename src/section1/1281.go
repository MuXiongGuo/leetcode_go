package main

func subtractProductAndSum(n int) int {
	s0, s1 := 1, 0
	for n > 0 {
		tmp := n % 10
		s0 *= tmp
		s1 += tmp
		n /= 10
	}
	return s0 - s1
}

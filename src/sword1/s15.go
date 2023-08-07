package main

func hammingWeight(num uint32) int {
	ans := 0
	for i := 0; i <= 31; i++ {
		x := num & (1 << i)
		if x != 0 {
			ans++
		}
	}
	return ans
}

// n&(n-1) 效果是消除掉最低为的1 不断消除 直到没1为止
func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}

//func main() {
//	fmt.Println(4 & 1)
//}

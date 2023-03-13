package main

// 哈希+前缀和 这次统计整除 多了一个数学小技巧 切记要取余统计 而不是加法枚举
// 负数取余 又是一个涉及数学的问题
func subarraysDivByK(nums []int, k int) int {
	ans, s, m := 0, 0, map[int]int{0: 1}
	for _, v := range nums {
		s += v
		modulus := (s%k + k) % k
		ans += m[modulus]
		m[modulus]++
	}
	return ans
}

// 更加数学 不过不难
func subarraysDivByK(nums []int, k int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range nums {
		sum += elem
		modulus := (sum%k + k) % k
		record[modulus]++
	}
	for _, cx := range record {
		ans += cx * (cx - 1) / 2
	}
	return ans
}

//func main() {
//	m := map[int]int{}
//	ans := m[8]
//	println(ans)
//}

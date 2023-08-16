package main

// 模拟 哈希
func circularGameLosers(n int, k int) []int {
	m := map[int]bool{}
	i := 0
	cur := 1
	for true {
		cur += i * k % n
		if cur > n {
			cur %= n
		}
		i++
		if m[cur] {
			break
		}
		m[cur] = true
	}
	ans := []int{}
	for j := 1; j <= n; j++ {
		if !m[j] {
			ans = append(ans, j)
		}
	}
	return ans
}
func main() {
	circularGameLosers(5, 3)
}

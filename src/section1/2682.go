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

// 官方更优雅
func circularGameLosers(n, k int) []int {
	visit := make([]bool, n)
	j := 0
	for i := k; !visit[j]; i += k {
		visit[j] = true
		j = (j + i) % n
	}
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	circularGameLosers(5, 3)
}

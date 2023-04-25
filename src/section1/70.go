package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 可以减少一维 只用有限个变量即可
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 可以减少一维 只用有限个变量即可
	r, p := 1, 2
	var q int
	for i := 3; i <= n; i++ {
		q = r + p
		r = p
		p = q
	}
	return q
}

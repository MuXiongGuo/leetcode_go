package main

// dp
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j != 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if i != 0 && j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 官方更简洁

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, g := range grid {
		for j, x := range g {
			if i > 0 {
				f[i][j] = max(f[i][j], f[i-1][j])
			}
			if j > 0 {
				f[i][j] = max(f[i][j], f[i][j-1])
			}
			f[i][j] += x
		}
	}
	return f[m-1][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 其实只用保存最后两行即可
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, g := range grid {
		pos := i % 2
		for j, x := range g {
			f[pos][j] = 0
			if i > 0 {
				f[pos][j] = max(f[pos][j], f[1-pos][j])
			}
			if j > 0 {
				f[pos][j] = max(f[pos][j], f[pos][j-1])
			}
			f[pos][j] += x
		}
	}
	return f[(m-1)%2][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

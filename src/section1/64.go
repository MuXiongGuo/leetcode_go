package main

// 暴力回溯 超时了
func dfs(grid [][]int, x, y, z int) int {
	if x == len(grid)-1 && y == len(grid[0])-1 {
		return z + grid[x][y]
	}
	if x == len(grid)-1 {
		return dfs(grid, x, y+1, z+grid[x][y])
	}
	if y == len(grid[0])-1 {
		return dfs(grid, x+1, y, z+grid[x][y])
	}
	right := dfs(grid, x, y+1, z+grid[x][y])
	down := dfs(grid, x+1, y, z+grid[x][y])
	if right < down {
		return right
	}
	return down
}

func minPathSum(grid [][]int) int {
	return dfs(grid, 0, 0, 0)
}

// 2. dp
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

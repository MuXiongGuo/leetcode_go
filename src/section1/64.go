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

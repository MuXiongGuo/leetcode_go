package main

import (
	"fmt"
)

func uniquePathsIII(grid [][]int) int {
	ans := 0
	cntZero := 0
	var x, y int

	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				x = i
				y = j
			} else if grid[i][j] == 0 {
				cntZero++
			}
		}
	}

	cntZero++
	grid[x][y] = 0
	var helper func(i, j int)
	helper = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if grid[i][j] == 2 {
			if cntZero == 0 {
				ans++
			}
			return
		}
		if grid[i][j] == 1 || grid[i][j] == -1 {
			return
		}
		// grid[i][j] == 0
		grid[i][j] = 1
		cntZero--
		helper(i+1, j)
		helper(i-1, j)
		helper(i, j+1)
		helper(i, j-1)
		grid[i][j] = 0
		cntZero++
		return

	}
	helper(x, y)
	return ans
}

func main() {
	grid := [][]int{{1, 0}, {2, 0}}
	x := uniquePathsIII(grid)
	fmt.Println(x)
}

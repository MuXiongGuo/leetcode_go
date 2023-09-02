package main

import (
	"fmt"
)

// 可以加一个空间优化，之前走到过的地方 不增加的话那么就在这个点上比较为false 之后走的时候 遇到false的直接结束即可
// 注意！！！不是简单的给每个节点标记false true  而是将这条路径 压缩成二进制 同样的路径的情况来判断能否通过
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

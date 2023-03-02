package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2)
	}
	for i, j := 1, 1; i != n-2 || j != n-1; {
		// get max
		val := -1
		val = max(val, grid[i][j])
		val = max(val, grid[i-1][j])
		val = max(val, grid[i+1][j])
		val = max(val, grid[i][j-1])
		val = max(val, grid[i][j+1])
		val = max(val, grid[i-1][j-1])
		val = max(val, grid[i+1][j+1])
		val = max(val, grid[i+1][j-1])
		val = max(val, grid[i-1][j+1])
		ans[i-1][j-1] = val
		// increase
		j++
		if j == n-1 && i != n-2 {
			j = 1
			i += 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 官方  很简洁 学习下
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 1; i < n-1; i++ {
		row := make([]int, n-2)
		for j := 1; j < n-1; j++ {
			mx := grid[i][j]
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if grid[r][c] > mx {
						mx = grid[r][c]
					}
				}
			}
			row[j-1] = mx
		}
		ans[i-1] = row
	}
	return ans
}

func main() {
	grid := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	ans := largestLocal(grid)
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans); j++ {
			print(ans[i][j], " ")
		}
		println()
	}
}

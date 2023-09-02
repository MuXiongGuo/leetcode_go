package main

// 回溯 但是超时了
//func minFallingPathSum(grid [][]int) int {
//	ans, tmp := math.MaxInt64, 0
//	n, m := len(grid), len(grid[0])
//	var helper func(i, j int)
//	helper = func(i, j int) {
//		if i == n {
//			ans = min(ans, tmp)
//			return
//		}
//		for x := 0; x < m; x++ {
//			if x == j {
//				continue
//			}
//			tmp += grid[i][x]
//			helper(i+1, x)
//			tmp -= grid[i][x]
//		}
//	}
//	helper(0, -1)
//	return ans
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

// 试试dp
// 不要用dp表示每一行 那样找不到递推规律
// 用dp[i][j] 表示第i行 选择grid[i][j]的最小值
// 最终答案是最后一行的min(dp[i][j])
// 每个格子一个dp 而不是一行一个!
// 附带时空优化
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}
	dp := make([]int, n)
	dp = grid[0]

	// 得到初始的minFirst和minSecond
	minFirst, minSecond := 0, 1
	if grid[0][minSecond] < grid[0][minFirst] {
		minFirst, minSecond = minSecond, minFirst
	}
	for i := 2; i < n; i++ {
		if grid[0][i] < grid[0][minFirst] {
			minSecond = minFirst
			minFirst = i
		} else if grid[0][i] < grid[0][minSecond] {
			minSecond = i
		}
	}
	// 开始
	for i := 1; i < n; i++ {
		tmpFirst, tmpSecond := -1, -1
		m1, m2 := dp[minFirst], dp[minSecond]
		for j := 0; j < n; j++ {
			if j == minFirst {
				dp[j] = grid[i][j] + m2
			} else {
				dp[j] = grid[i][j] + m1
			}
			if tmpFirst == -1 {
				tmpFirst = j
			} else if tmpSecond == -1 {
				tmpSecond = j
				if dp[tmpSecond] < dp[tmpFirst] {
					tmpFirst, tmpSecond = tmpSecond, tmpFirst
				}
			} else if dp[j] < dp[tmpFirst] {
				tmpSecond = tmpFirst
				tmpFirst = j
			} else if dp[j] < dp[tmpSecond] {
				tmpSecond = j
			}
		}
		// 更新minFirst和minSecond
		minFirst, minSecond = tmpFirst, tmpSecond
	}
	return dp[minFirst]
}
func main() {
	minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

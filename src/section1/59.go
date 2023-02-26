package main

func generateMatrix(n int) [][]int {
	// 构建二维基本步骤 很简单
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	start, end, val := 0, n-1, 1
	for start <= end {
		// 1
		for i := start; i <= end; i++ {
			ans[start][i] = val
			val++
		}
		// 2
		for i := start + 1; i <= end; i++ {
			ans[i][end] = val
			val++
		}
		// 3
		for i := end - 1; i >= start; i-- {
			ans[end][i] = val
			val++
		}
		// 4
		for i := end - 1; i >= start+1; i-- {
			ans[i][start] = val
			val++
		}
		start++
		end--
	}
	return ans
}

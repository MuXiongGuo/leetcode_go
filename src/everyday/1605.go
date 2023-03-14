package main

// 数学而已  构建一个模拟逻辑即可
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	// 枚举m 弄个n的求和
	for i := 0; i < m; i++ {
		curSum := colSum[i]
		for j := 0; j < n; j++ {
			// 竖着的或者横着的 用光了
			if rowSum[j] == 0 || curSum == 0 {
				ans[j][i] = 0
			} else if curSum >= rowSum[j] {
				ans[j][i] = rowSum[j]
				curSum -= rowSum[j]
				rowSum[j] = 0
			} else {
				ans[j][i] = curSum
				rowSum[j] -= curSum
				curSum = 0
			}
		}
	}
	return ans
}

// 灵神优化 类似蛇形向右向下走填充 其余全为0 忽略创建矩阵时间的话只花了O(m+n)
func restoreMatrix2(rowSum []int, colSum []int) [][]int {
	ans := make([][]int, len(rowSum))
	for i := range ans {
		ans[i] = make([]int, len(colSum))
	}
	for i, j := 0, 0; i < len(rowSum) && j < len(colSum); {
		rs, cs := rowSum[i], colSum[j]
		if rs >= cs {
			ans[i][j] = cs
			rowSum[i] -= cs
			j++
		} else {
			ans[i][j] = rs
			colSum[j] -= rs
			i++
		}
	}
	return ans
}

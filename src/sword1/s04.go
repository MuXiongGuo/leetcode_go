package main

// 以前算法导论 刷到过 很有意思 保证每个方向是固定加减 从右上或者左下角开始
func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 特殊情况 要考虑
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	for i, j := 0, m-1; i < n && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

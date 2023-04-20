package main

// 回溯 标记一下用过的即可
func helper(board [][]byte, word string, used [][]bool, x, y, z int) bool {
	if y >= len(board[0]) || x >= len(board) || y < 0 || x < 0 || used[x][y] == true || board[x][y] != word[z] {
		return false
	}
	if z == len(word)-1 {
		return true
	}
	used[x][y] = true
	ans := helper(board, word, used, x+1, y, z+1) || helper(board, word, used, x-1, y, z+1) || helper(board, word, used, x, y+1, z+1) || helper(board, word, used, x, y-1, z+1)
	used[x][y] = false
	return ans
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				ans := helper(board, word, used, i, j, 0)
				if ans == true {
					return true
				}
			}
		}
	}
	return false
}

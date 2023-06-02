package main

// 回溯
func generateParenthesis(n int) []string {
	var Ans []string
	var dfs func(n, count, cur int, s string)
	// (+1  )-1
	dfs = func(n, count, cur int, s string) {
		if cur < 0 {
			return
		}
		if count == n {
			if cur == 0 {
				Ans = append(Ans, s)
			}
			return
		}
		dfs(n, count+1, cur+1, s+"(")
		dfs(n, count+1, cur-1, s+")")
	}

	dfs(2*n, 0, 0, "")
	return Ans
}

//func main() {
//	fmt.Println(generateParenthesis(1))
//}

//
//func main() {
//	A := make([]int, 4, 10)
//	A[0] = 1
//	A[1] = 2
//	A[2] = 3
//	A[3] = 4
//
//预留足够的话就不可 但不预留就可！
//	B := append(A, 9)
//	c := append(A, 6)
//
//	fmt.Println(A)
//	fmt.Println(B)
//	fmt.Println(c)
//
//	s := "da"
//	z := s+
//}

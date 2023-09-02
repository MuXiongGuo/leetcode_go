package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	// 并查集模板
	fa := make([]int, 26)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from = find(from)
		to = find(to)
		if from != to {
			fa[from] = to
		}
	}
	// 97-122
	for i := 97; i <= 122; i++ {
		if i == 'b' || i == 'q' || i == 'd' || i == 'p' {
			merge(i-97, 'b'-97)
		} else if i == 'n' || i == 'u' {
			merge(i-97, 'n'-97)
		}
	}

	for i := 0; i < n; i++ {
		if helper(s[i], fa) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func helper(s string, fa []int) bool {
	n := len(s)
	i, j := 0, n-1
	for i < j {
		el1, el2 := s[i], s[j]
		if fa[el1-97] != fa[el2-97] {
			// 测试是否能分割
			if el1 == 'w' {
				if s[j] == 'v' && s[j-1] == 'v' {
					j -= 2
					i++
					continue
				}
			}
			if el2 == 'w' {
				if s[i] == 'v' && s[i+1] == 'v' {
					i += 2
					j--
					continue
				}
			}

			if el1 == 'm' {
				if fa[s[j]-97] == 'n'-97 && fa[s[j-1]-97] == 'n'-97 {
					j -= 2
					i++
					continue
				}
			}
			if el2 == 'm' {
				if fa[s[i]-97] == 'n'-97 && fa[s[i+1]-97] == 'n'-97 {
					i += 2
					j--
					continue
				}
			}
			return false
		}
		i++
		j--
	}
	return true
}

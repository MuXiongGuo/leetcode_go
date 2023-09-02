package main

package main

func letterCombinations(digits string) []string {
	n := len(digits)
	ans := []string{}
	path := make([]byte, n)
	m := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var helper func(i int)
	helper = func(x int) {
		if x == n {
			ans = append(ans, string(path))
			return
		}
		s := m[digits[x]]
		for i := 0; i < len(s); i++ {
			path[x] = s[i]
			helper(x + 1)
		}
	}
	if n != 0 {
		helper(0)
	}
	return ans
}

func main() {
	digits := "23"
	println(letterCombinations(digits))
}

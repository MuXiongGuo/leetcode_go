package main

import "strconv"

func getPermutation(n int, k int) string {
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i + 1
	}
	for j := 0; j < n; j++ {
		if k <= 1 {
			break
		}
		if count(n-j-1)+1 > k {
			continue
		}
		x := k / count(n-j-1)
		k -= x * count(n-j-1)
		ans[j], ans[j+x] = ans[j+x], ans[j]
		if k <= 1 {
			break
		}
	}
	res := ""
	for _, el := range ans {
		res += strconv.Itoa(el)
	}
	return res
}

func count(n int) int {
	if n == 1 {
		return 2
	}

	ans := 1
	for n > 0 {
		ans *= n
		n--
	}
	return ans
}

func main() {
	getPermutation(5, 2)
}

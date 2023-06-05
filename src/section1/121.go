package main

func maxProfit(prices []int) int {
	cur := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > cur {
			ans = max(ans, prices[i]-cur)
		} else {
			cur = prices[i]
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

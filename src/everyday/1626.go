package main

import "sort"

// dp 最长子数组问题, 按照年龄和分数排序都可以滴
func bestTeamScore(scores []int, ages []int) (ans int) {
	type pair struct{ age, score int }
	n := len(scores)
	var pairSlice []pair
	for i := 0; i < n; i++ {
		pairSlice = append(pairSlice, pair{
			age:   ages[i],
			score: scores[i],
		})
	}
	sort.Slice(pairSlice, func(i, j int) bool {
		a, b := pairSlice[i], pairSlice[j]
		return a.age < b.age || a.age == b.age && a.score < b.score // 直接省略一个if判断了 细节优雅
	})
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if pairSlice[i].score >= pairSlice[j].score {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += pairSlice[i].score
		ans = max(ans, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 灵神的  优雅
func bestTeamScore(scores, ages []int) (ans int) {
	n := len(scores)
	type pair struct{ score, age int }
	a := make([]pair, n)
	for i, s := range scores {
		a[i] = pair{s, ages[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.score < b.score || a.score == b.score && a.age < b.age
	})

	f := make([]int, n)
	for i, p := range a {
		for j, q := range a[:i] {
			if q.age <= p.age {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += p.score
		ans = max(ans, f[i])
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

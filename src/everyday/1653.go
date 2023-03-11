package main

import "strings"

// 1.前后缀分解 既然找不到规律 不如考虑枚举所有可能的最终结果
// 不断构造一条分割中线 因为所有的情况一定是可知道的 并不是很复杂的
func minimumDeletions(s string) int {
	a, b, ans := 0, 0, 0
	for _, v := range s {
		if v == 'a' {
			a++
		}
	}
	ans = a
	for _, v := range s {
		if v == 'b' {
			b++
		} else {
			a--
		}
		ans = min(ans, a+b)
	}
	return ans
}

// 参考答案
func minimumDeletions(s string) int {
	del := strings.Count(s, "a")
	ans := del
	for _, c := range s {
		// 'a' -> -1    'b' -> 1  不用if else 会大大提高运行速度
		del += int((c-'a')*2 - 1)
		if del < ans {
			ans = del
		}
	}
	return ans
}

// 2.动态规划  分析后一状态与前一状态 重点是得到转移方程
// 若为b 则后与前一样 若为a 则面临两种选择 以此得到状态方程
// b:d(n)=d(n+1)
// a:d(n)=min(d(n-1)+1, cntb)
// 之前没找到状态就是忽略了cntb
func minimumDeletions(s string) int {
	d, b := 0, 0
	for _, v := range s {
		if v == 'a' {
			d = min(d+1, b)
		} else {
			b++
		}
	}
	return d
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

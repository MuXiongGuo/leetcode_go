package main

import "math"

// 前后缀
func minimumAverageDifference(nums []int) int {
	sum, n := 0, len(nums)
	mIndex, mVal := 0, math.MaxInt64
	for _, v := range nums {
		sum += v
	}
	cnt, s := 0, 0
	for i, v := range nums {
		cnt++
		s += v
		n--
		sum -= v
		var k int
		if n == 0 {
			k = int(math.Abs(float64(s / cnt)))
		} else {
			k = int(math.Abs(float64((s / cnt) - (sum / n))))
		}
		if mVal < k {
			mVal = k
			mIndex = i
		}
	}
	return mIndex
}

// 灵神 利用切片 更优雅
func minimumAverageDifference(nums []int) (ans int) {
	pre, suf, n := 0, 0, len(nums)
	for _, v := range nums {
		suf += v
	} // 后缀和
	minDiff := math.MaxInt64
	for i, v := range nums[:n-1] {
		pre += v // 前缀和
		suf -= v // 后缀和
		d := abs(pre/(i+1) - suf/(n-1-i))
		if d < minDiff {
			minDiff, ans = d, i
		}
	}
	if (pre+nums[n-1])/n < minDiff {
		ans = n - 1
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	minimumAverageDifference([]int{2, 5, 3, 9, 5, 3})
}

package main

import (
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	s := make([][2]int, n)
	for i := range s {
		s[i][0] = dist[i]
		s[i][1] = speed[i]
	}
	sort.Slice(s, func(i, j int) bool {
		return float64(s[i][0])/float64(s[i][1]) < float64(s[j][0])/float64(s[j][1])
	})
	ans := 0
	for time := 0; time < n; time++ {
		if s[time][0]-time*s[time][1] <= 0 {
			break
		}
		ans++
	}
	return ans
}

// 官方 直接得到到达时间 因为我方的攻击时间为整数，因此可以将怪物到达时间向上取整，可以达到避免浮点数误差的效果
// 然后直接比较即可 没必要记录谁是谁
func eliminateMaximum2(dist []int, speed []int) int {
	n := len(dist)
	arrivalTimes := make([]int, n)
	for i := 0; i < n; i++ {
		arrivalTimes[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(arrivalTimes)
	for i := 0; i < n; i++ {
		if arrivalTimes[i] <= i {
			return i
		}
	}
	return n
}

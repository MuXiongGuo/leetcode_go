package main

// 回溯 每一步选择一种笔 但是会有重复多余情况  不如直接使用枚举来做
//func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
//	ans := 1
//	var helper func(total int)
//	helper = func(total int) {
//		if total < cost1 && total < cost2 {
//			return
//		}
//		if total >= cost1 {
//			helper(total - cost1)
//		}
//		if total >= cost2 {
//			ans++
//			helper(total - cost2)
//		}
//	}
//	helper(total)
//	return int64(ans)
//}

// 暴力
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	ans := 0
	for i := 0; i*cost1 <= total; i++ {
		ans += (total-i*cost1)/cost2 + 1
	}
	return int64(ans)
}

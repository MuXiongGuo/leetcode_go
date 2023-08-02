package main

func flipgame(fronts []int, backs []int) int {
	doubleSide := map[int]bool{}
	n := len(fronts)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			doubleSide[fronts[i]] = true
		}
	}
	res := 2001
	for i := 0; i < n; i++ {
		v0, v1 := fronts[i], backs[i]
		if doubleSide[v0] && doubleSide[v1] {
			continue
		}
		if doubleSide[v0] || doubleSide[v1] {
			var v2 int
			if doubleSide[v0] {
				v2 = v1
			} else {
				v2 = v0
			}
			res = min(res, v2)
			continue
		}
		res = min(res, min(v0, v1))
	}
	if res == 2001 {
		return 0
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//
//func main() {
//	m := map[int]int{}
//	if m[3] == 4 {
//		fmt.Println(m)
//	}
//	fmt.Println(len(m))
//	fmt.Println(m[88])
//
//	m2 := map[int]struct{}{}
//	m2[8] = struct{}{}
//	fmt.Println(m2[8])
//	fmt.Println(m2[3])
//	if _, ok := m2[8]; ok {
//		fmt.Println("hello")
//	}
//
//	m3 := map[int]bool{}
//
//	fmt.Println(m3[444])
//}

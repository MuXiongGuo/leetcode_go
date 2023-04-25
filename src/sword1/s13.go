package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func movingCount(m int, n int, k int) int {
	// 并查集模板
	fa := make([]int, m*n)
	size := make([]int, m*n)
	for i := range fa {
		fa[i] = i
		size[i] = 1
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
			size[to] += size[from]
		}
	}

	// 开始操作
	countValue := func(x, y int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x /= 10
		}
		for y != 0 {
			res += y % 10
			y /= 10
		}
		return res
	}
	// 不用排序权值  直接扫一遍 连接起来
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if countValue(i, j) <= k {
				// 连接周围四个点
				for _, d := range dirs {
					x, y := i+d.x, j+d.y
					if x >= 0 && x < m && y >= 0 && y < n && countValue(x, y) <= k {
						merge(x*n+y, i*n+j)
					}
				}
			}
		}
	}

	return size[fa[0]]
}

func main() {
	//// 并查集模板
	//fa := make([]int, n)
	//size := make([]int, n)
	//var find func(int) int
	//find = func(x int) int {
	//	if fa[x] != x {
	//		fa[x] = find(fa[x])
	//	}
	//	return fa[x]
	//}
	//merge := func(from, to int) {
	//	from = find(from)
	//	to = find(to)
	//	if from != to {
	//		fa[from] = to
	//		size[to] += size[from]
	//	}
	//}
	//
	//// 矩阵元素从小到大排序，小模板
	//type tuple struct{ x, i, j int }
	//a := make([]tuple, 0, mn)
	//for i, row := range grid {
	//	for j, x := range row {
	//		a = append(a, tuple{x, i, j})
	//	}
	//}
	//sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	movingCount(2, 3, 1)
}

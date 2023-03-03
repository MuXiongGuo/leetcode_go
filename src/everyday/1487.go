package main

import (
	"fmt"
	"strconv"
)

// 时间复杂度分析 额外的else循环最多n次  时间复杂度是 names的元素数量与长度的累加
func getFolderNames(names []string) []string {
	m := map[string]int{}
	ans := make([]string, len(names))
	count := 0
	for _, name := range names {
		n, ok := m[name]
		if !ok {
			m[name] = 1
			ans[count] = name
		} else {
			for _, y := m[name+"("+strconv.Itoa(n)+")"]; y; {
				n++
				_, y = m[name+"("+strconv.Itoa(n)+")"]
			}
			m[name] = n + 1 // 细节 是n+1哈
			m[name+"("+strconv.Itoa(n)+")"] = 1
			ans[count] = name + "(" + strconv.Itoa(n) + ")"
		}
		count++
	}
	return ans
}

// 官方 思路和我的一样 更简洁一些, range的index 以及判断存在时 直接用m[...] > 0
// 而且直接用了原来的name 没有新开辟数组
func getFolderNames(names []string) []string {
	d := map[string]int{}
	for i, name := range names {
		if k, ok := d[name]; ok {
			for {
				newName := fmt.Sprintf("%s(%d)", name, k)
				if d[newName] == 0 {
					d[name] = k + 1
					names[i] = newName
					break
				}
				k++
			}
		}
		d[names[i]] = 1
	}
	return names
}
func getFolderNames(names []string) []string {
	ans := make([]string, len(names))
	index := map[string]int{}
	for p, name := range names {
		i := index[name]
		if i == 0 {
			index[name] = 1
			ans[p] = name
			continue
		}
		for index[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		t := name + "(" + strconv.Itoa(i) + ")"
		ans[p] = t
		index[name] = i + 1
		index[t] = 1
	}
	return ans
}

package main

// 1.字母存储 一般不用map   可以用数组 m := [26]int
func firstUniqChar(s string) byte {
	m := map[byte]int{}
	for _, c := range s {
		m[byte(c)]++
	}
	for _, c := range s {
		if m[byte(c)] == 1 {
			return byte(c)
		}
	}
	return ' '
}

// 2.优化 第二次遍历其实多少有些浪费时间  可以s+map 大大节约了时间
// 若遇到重复的就标记为-1  只有一次的标记为index 而不是统计次数
// 第二次 遍历map 选择最小index即可
func firstUniqChar(s string) byte {
	m := [26]int{}
	for i, v := range s {
		if m[v-'a'] == 0 {
			m[v-'a'] = i + 1
		} else {
			m[v-'a'] = -1
		}
	}
	ans := 50001
	for _, i := range m {
		if i != -1 && i != 0 && i < ans {
			ans = i
		}
	}
	if ans == 50001 {
		return ' '
	}
	return s[ans-1]
}

// 用队列做  golang的切片 可以模拟队列  从头部删除毕竟底层不用管
type pair struct {
	ch  byte
	pos int
}

func firstUniqChar(s string) byte {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	q := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return s[q[0].pos]
	}
	return ' '
}

package main

import "strings"

type StreamChecker struct {
	m map[string]struct{}
	s strings.Builder
}

func Constructor(words []string) StreamChecker {
	ans := StreamChecker{m: map[string]struct{}{}}
	for _, v := range words {
		ans.m[v] = struct{}{}
	}
	return ans
}

func (this *StreamChecker) Query(letter byte) bool {
	this.s.WriteByte(letter)
	for i := len(this.s.String()) - 1; i >= 0; i-- {
		if _, ok := this.m[this.s.String()[i:]]; ok {
			return true
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

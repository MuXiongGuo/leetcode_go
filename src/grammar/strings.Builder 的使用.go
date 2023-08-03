package main

import (
	"fmt"
	"strings"
)

func main() {
	// 直接用string + 性能差 不会智能的扩容
	// func (b *Builder) Write(p []byte) (int, error)
	// func (b *Builder) WriteByte(c byte) error
	// func (b *Builder) WriteRune(r rune) (int, error)
	// func (b *Builder) WriteString(s string) (int, error)

	var b strings.Builder
	b.WriteString("abc")
	b.WriteString("def")
	fmt.Println(b.String())
}

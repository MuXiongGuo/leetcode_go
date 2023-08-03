package main

import "strings"

func removeComments(source []string) []string {
	var ans, all strings.Builder
	for _, el := range source {
		all.WriteString(el)
		all.WriteString("\n")
	}
	type pair struct {
		a, b int
	}
	var allIndex []pair
	for i := 0; i < len(all.String())-1; i++ {
		if all.String()[i:i+2] == "//" {
			tmp := pair{a: i}
			for j := i + 2; j < len(all.String())-1; j++ {
				if all.String()[j:j+2] == "\n" {
					tmp.b = j + 1
				}
			}
		} else if all.String()[i:i+2] == "/*" {
			tmp := pair{a: i}
			for j := i + 2; j < len(all.String())-1; j++ {
				if all.String()[j:j+2] == "\n" {
					tmp.b = j + 1
				}
			}
		}
	}
}

//func main() {
//	s := "abc"
//	fmt.Println(s[1:3])
//
//}

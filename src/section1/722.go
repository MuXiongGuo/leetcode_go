package main

//func removeComments(source []string) []string {
//	var ans, all strings.Builder
//	for _, el := range source {
//		all.WriteString(el)
//		all.WriteString("\n")
//	}
//	type pair struct {
//		a, b int
//	}
//	var allIndex []pair
//	for i := 0; i < len(all.String())-1; i++ {
//		if all.String()[i:i+2] == "//" || all.String()[i:i+2] == "/*" {
//			tmp := pair{a: i}
//			flagStar := true
//			if all.String()[i:i+2] == "//" {
//				flagStar = false
//			}
//			for j := i + 2; j < len(all.String())-1; j++ {
//				if all.String()[j:j+2] == "*/" && flagStar {
//					tmp.b = j + 2
//					i = j + 1
//					break
//				} else if all.String()[j:j+1] == "\n" && !flagStar {
//					tmp.b = j
//					i = j + 1
//					break
//				}
//			}
//			allIndex = append(allIndex, tmp)
//		}
//	}
//	for i := 0; i <= len(allIndex); i++ {
//		var start, end int
//		if i == 0 {
//			start, end = 0, allIndex[i].a
//		} else if i == len(allIndex) {
//			start, end = allIndex[i-1].b+1, len(all.String())
//		} else {
//			start, end = allIndex[i-1].b+1, allIndex[i].a
//		}
//		ans.WriteString(all.String()[start:end])
//	}
//	xxx := strings.Split(ans.String(), "\n")
//	return xxx[:len(xxx)-1]
//}

func main() {
	source := []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}
	removeComments(source)
	//var all strings.Builder
	//for _, el := range source {
	//	all.WriteString(el)
	//	all.WriteString("\n")
	//}
	//type pair struct {
	//	a, b int
	//}
	//var allIndex []pair
	//for i := 0; i < len(all.String())-1; i++ {
	//	if all.String()[i:i+2] == "//" || all.String()[i:i+2] == "/*" {
	//		tmp := pair{a: i}
	//		flagStar := true
	//		if all.String()[i:i+2] == "//" {
	//			flagStar = false
	//		}
	//		for j := i + 2; j < len(all.String())-1; j++ {
	//			if all.String()[j:j+2] == "*/" && flagStar {
	//				tmp.b = j + 2
	//				i = j + 1
	//				break
	//			} else if all.String()[j:j+1] == "\n" && !flagStar{
	//				tmp.b = j
	//				i = j + 1
	//				break
	//			}
	//		}
	//		allIndex = append(allIndex, tmp)
	//	}
	//}
	//fmt.Println(all.String()[72:89])
}

//	func main() {
//		b := strings.Builder{}
//		b.WriteString("123")
//		b.WriteString("\n")
//		fmt.Println(b)
//	}

func removeComments(source []string) []string {
	res := []string{}
	new_line := []byte{}
	in_block := false
	for _, line := range source {
		for i := 0; i < len(line); i++ {
			if in_block {
				if i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
					in_block = false
					i++
				}
			} else {
				if i+1 < len(line) && line[i] == '/' && line[i+1] == '*' {
					in_block = true
					i++
				} else if i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
					break
				} else {
					new_line = append(new_line, line[i])
				}
			}
		}
		if !in_block && len(new_line) > 0 {
			res = append(res, string(new_line))
			new_line = []byte{}
		}
	}
	return res
}

// python  正则 灵神
//class Solution:
//def removeComments(self, source: List[str]) -> List[str]:
//# 匹配所有 // 和 /* */，后者用非贪婪模式。将所有匹配结果替换成空串。最后移除多余空行。
//return list(filter(None, re.sub('//.*|/\*(.|\n)*?\*/', '', '\n'.join(source)).split('\n')))

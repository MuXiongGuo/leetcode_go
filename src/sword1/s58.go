package main

func reverseLeftWords(s string, n int) string {
	s1 := []byte(s)
	reverseString(s1[:n])
	reverseString(s1[n:])
	reverseString(s1)
	return string(s1)
}

func reverseString(s []byte) {
	low, high := 0, len(s)-1
	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
}
func main() {
	s := "abcdefg"
	n := 2
	println(reverseLeftWords(s, n))
}

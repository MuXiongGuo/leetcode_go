package main

func replaceSpace(s string) string {
	var ans string
	for _, v := range s {
		if v == ' ' {
			ans += "%20"
		} else {
			ans += string(v)
		}
	}
	return ans
}

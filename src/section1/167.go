package main

func twoSum(numbers []int, target int) []int {
	ans := []int{}
	l, r := 0, len(numbers)-1
	for l < r {
		cur := numbers[l] + numbers[r]
		if cur < target {
			l++
		} else if cur > target {
			r--
		} else {
			ans = append(ans, numbers[l], numbers[r])
			return ans
		}

	}
	return ans
}

package main

func removeElement(nums []int, val int) int {
	ans := 0
	for _, element := range nums {
		if element != val {
			nums[ans] = element
			ans += 1
		}
	}
	return ans
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
}

package main

// 异或运算，数字与数字本身异或结果为0 与0异或结果为本身 满足结合律分配律
// x^x = 0  x^0 = x
// 所有数字异或即可
func singleNumber(nums []int) int {
	s := nums[0]
	for i := 1; i < len(nums); i++ {
		s ^= nums[i]
	}
	return s
}

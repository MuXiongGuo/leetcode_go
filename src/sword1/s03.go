package main

// 哈希 线性空间
func findRepeatNumber(nums []int) int {
	m := make([]int, len(nums))
	for _, v := range nums {
		if m[v] == 1 {
			return v
		} else {
			m[v] = 1
		}
	}
	return -1
}

// 原地交换 利用数据特点 常数空间完成  每个数都有唯一位置了
func findRepeatNumber2(nums []int) int {
	// range 不是动态的 坑人
	for i := 0; i < len(nums); {
		if nums[i] == nums[nums[i]] && i != nums[i] {
			return nums[i]
		}
		if nums[i] == nums[nums[i]] {
			i++
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// 更有逻辑
func findRepeatNumber3(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
		} else if nums[i] == nums[nums[i]] {
			return nums[i]
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

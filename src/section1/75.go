package main

func sortColors(nums []int) {
	// 线性 原地排序 但是 两趟
	n0, n1, n2 := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			n0++
		case 1:
			n1++
		case 2:
			n2++
		}
	}
	for i := 0; i < n0; i++ {
		nums[i] = 0
	}
	for i := n0; i < n0+n1; i++ {
		nums[i] = 1
	}
	for i := n0 + n1; i < n0+n1+n2; i++ {
		nums[i] = 2
	}
}

func sortColors(nums []int) {
	// 线性 原地排序 只扫描一次！！！
	// 设置一个左右指针 很简单
	left, right, n := 0, len(nums)-1, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			break
		} else {
			left++
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] != 2 {
			break
		} else {
			right--
		}
	}

	Swap0 := func(i int) {
		nums[left], nums[i] = nums[i], nums[left]
		for left < n || nums[left] != 0 {
			left++
		}
	}
	Swap2 := func(i int) {
		nums[right], nums[i] = nums[i], nums[right]
		for right >= 0 || nums[right] != 2 {
			right--
		}
	}

	for p := left; p <= right; p++ {
		if nums[p] == 0 {
			Swap0(p)
			if nums[p] == 2 {
				Swap2(p)
			}
		} else if nums[p] == 2 {
			Swap2(p)
			if nums[p] == 0 {
				Swap0(p)
			}
		}
	}

}

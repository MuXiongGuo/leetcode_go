package main

// O(n) space use another array to store
func rotate(nums []int, k int) {
	n := len(nums)
	numsCopy := make([]int, n)
	copy(numsCopy, nums) // copy 函数要保证两个数组一样大，否则只能copy min个资源
	k = k % n
	y := 0
	for id := range nums {
		if k > 0 {
			nums[id] = numsCopy[n-k]
			k--

		} else {
			nums[id] = numsCopy[y]
			y++
		}
	}
}

// O(n)空间 use append copy
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	n1, n2 := make([]int, k), make([]int, n-k)
	copy(n1, nums[n-k:])
	copy(n2, nums[:n-k+1])
	//nums = append(n1, n2...) // leetcode里过不了 因为它检查底层数组而不是检查nums这个结构体
	i := 0
	for _, el := range n1 {
		nums[i] = el
		i++
	}
	for _, el := range n2 {
		nums[i] = el
		i++
	}
}

// O(1)空间
func rotate3(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverseSlice(nums)
	reverseSlice(nums[:k])
	reverseSlice(nums[k:])
}

func reverseSlice(s []int) {
	n := len(s) - 1
	i := 0
	for i < n {
		s[i], s[n] = s[n], s[i]
		i++
		n--
	}
}

// 若不用翻转3次的方法，那么可得到规律  旧的在位置i处的元素 应该 放在(i+k) mod n 处这样有这个公式就可以就地替换了
// 不好理解
func rotate4(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

// 单变量的, 最后一定会回来, tmp 保存 nums[0]元素  nums[0]本该在nums[x] 交换tmp和nums[x] 此时x处继续操作即可
func rotate5(nums []int, k int) {
	n := len(nums)
	k %= n
	count := 0
	for start := 0; start < n && count < n; start++ {
		pre, tmp := start, nums[start]
		for {
			x := (pre + k) % n
			tmp, nums[x] = nums[x], tmp
			pre = x
			count++
			if start == x {
				break
			}
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	rotate5([]int{-1, -100, 3, 99}, 2)
}

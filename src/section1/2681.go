package main

import "sort"

// 1.先排序 再观察规律
// 2.我们按照如下规律 求每个以nums[i]结尾的和，所以的加起来就是结果
// 3.而每个nums[i]的数值和为 nums[i]*nums[i](a0+a1+...ai-1)要求出前面的最小值
// 4.利用前缀和可以减少O(n)  不用再求(a0+a1+...ai-1) 有前缀和的话
// 5.利用动态规划吗，因为中间步骤较麻烦， 设dp[i]为以nums[i]结尾的所有的最小值之和， dp[i] = nums[i]+dp[j](0<=j<i)
// 6.前缀和模板 preSum[0] = 0     preSum[i+1] = preSum[i]+dp[i]
// 7.dp[i]=nums[i]+pre_sum[i] 最终优化成这样
// 8.小心 ans 和 dp preSum 都可能溢出

// 错误示范 以为只是简单的连续的  其实不连续要用dp
//func sumOfPower(nums []int) int {
//	ans := 0
//	mod := int(1e9 + 7)
//	preSum := make([]int, len(nums)+1)
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] < nums[j]
//	})
//	for i := 1; i <= len(nums); i++ {
//		preSum[i] = preSum[i-1] + nums[i-1]
//	}
//	for i := 0; i < len(nums); i++ {
//		ans += nums[i] * nums[i] * preSum[i+1]
//		ans = ans % mod
//	}
//	return ans
//}

// 其实preSum与dp只与前一个状态有关 可以优化
func sumOfPower(nums []int) int {
	n := len(nums)
	ans := 0
	mod := int(1e9 + 7)
	preSum := make([]int, n+1)
	dp := make([]int, n)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < n; i++ {
		dp[i] = (nums[i] + preSum[i]) % mod
		preSum[i+1] = (preSum[i] + dp[i]) % mod
		ans = (ans + (nums[i]*nums[i]%mod*dp[i])%mod) % mod // 每一步都可能溢出
	}
	return ans
}

// 常数空间版
func sumOfPower(nums []int) int {
	n := len(nums)
	ans := 0
	mod := int(1e9 + 7)
	preSum := 0
	dp := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < n; i++ {
		dp = (nums[i] + preSum) % mod
		preSum = (preSum + dp) % mod
		ans = (ans + (nums[i]*nums[i]%mod*dp)%mod) % mod // 每一步都可能溢出
	}
	return ans
}

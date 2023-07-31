package main

import "sort"

// 利用153题 不够优雅
func search(nums []int, target int) int {
	n := len(nums)
	minIdex := findMin(nums)
	var ans int
	if target == nums[minIdex] {
		return minIdex
	} else if target > nums[minIdex] && nums[n-1] >= nums[0] {
		ans = sort.SearchInts(nums, target)
	} else if target > nums[minIdex] && nums[n-1] < nums[0] {
		if target <= nums[n-1] {
			ans = sort.SearchInts(nums[minIdex:], target) + minIdex
		} else if target <= nums[minIdex-1] {
			ans = sort.SearchInts(nums[:minIdex], target)
		}
	}
	if ans < len(nums) && nums[ans] == target {
		return ans
	}
	return -1
}

// 153题拿来用
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2 // 防止溢出
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return low
}

// c++  很巧妙
//class Solution {
//public:
//    int search(vector<int>& nums, int target) {
//        int n = (int)nums.size();
//        if (!n) {
//            return -1;
//        }
//        if (n == 1) {
//            return nums[0] == target ? 0 : -1;
//        }
//        int l = 0, r = n - 1;
//        while (l <= r) {
//            int mid = (l + r) / 2;
//            if (nums[mid] == target) return mid;
//            if (nums[0] <= nums[mid]) {
//                if (nums[0] <= target && target < nums[mid]) {
//                    r = mid - 1;
//                } else {
//                    l = mid + 1;
//                }
//            } else {
//                if (nums[mid] < target && target <= nums[n - 1]) {
//                    l = mid + 1;
//                } else {
//                    r = mid - 1;
//                }
//            }
//        }
//        return -1;
//    }
//};
//

package main

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。

进阶：你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
- 0 <= nums.length <= 105
- -10^9<= nums[i]<= 10^9
- nums是一个非递减数组
- -10^9<= target<= 10^9

链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
*/
var notFound = []int{-1, -1}

func searchRange(nums []int, target int) []int {
	n := len(nums)

	lo, hi := 0, n-1
	mid := -1
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			break
		}
	}

	if mid == -1 || nums[mid] != target {
		return notFound
	}

	left, right := mid, mid
	for left >= 0 {
		if nums[left] != target {
			break
		}
		left--
	}
	for right < n {
		if nums[right] != target {
			break
		}
		right++
	}

	return []int{left + 1, right - 1}
}

var testCases = []struct {
	nums   []int
	target int
	result []int
}{
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{1}, 1, []int{0, 0}},
}

func main() {
	for _, tc := range testCases {
		got := searchRange(tc.nums, tc.target)
		want := tc.result
		ok := true
		for i := range want {
			if got[i] != want[i] {
				ok = false
				break
			}
		}

		if ok {
			fmt.Println("ok")
		} else {
			fmt.Printf("got: %d, want: %d\n", got, want)
		}
	}

}

package main

import (
	"fmt"
)

/*
55. 跳跃游戏

给定一个非负整数数组nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度（注意：最大！）。
判断你是否能够到达最后一个下标。

示例1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

示例2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

提示：
- 1 <= nums.length <= 3 * 10^4
- 0 <= nums[i] <= 10^5

链接：https://leetcode-cn.com/problems/jump-game
*/
// 检查是否有0元素，如果没有，则一定能到达终点
// 如果有，则看能否跳过这个元素
func canJump(nums []int) bool {
	n := len(nums)

	if n == 1 {
		return true
	}

	// 找0，最后一个是否为0不重要
	endIdx := n - 1
	i := 0
	for i < endIdx {
		if nums[i] == 0 {
			if !canReach(nums, i-1, i+1) {
				return false
			}
		}
		i++
	}

	return true
}

// 从[0, end]的范围内，能否到达target
func canReach(nums []int, end int, target int) bool {
	for i := end; i >= 0; i-- {
		if nums[i] >= target-i {
			return true
		}
	}

	return false
}

var testCases = []struct {
	nums   []int
	result bool
}{
	{[]int{0}, true},
	{[]int{0, 1}, false},
}

func main() {
	for _, tc := range testCases {
		got := canJump(tc.nums)
		want := tc.result
		if got != want {
			fmt.Printf("got: %v, want: %v\n", got, want)
		} else {
			fmt.Println("ok")
		}
	}

}

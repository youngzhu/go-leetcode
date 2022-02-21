package main

import (
	"fmt"
	v1 "github.com/youngzhu/go-leetcode/p055/v1"
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
var testCases = []struct {
	nums   []int
	result bool
}{
	{[]int{0}, true},
	{[]int{0, 1}, false},
}

func main() {
	for _, tc := range testCases {
		got := v1.CanJump(tc.nums)
		want := tc.result
		if got != want {
			fmt.Printf("got: %v, want: %v\n", got, want)
		} else {
			fmt.Println("ok")
		}
	}

}

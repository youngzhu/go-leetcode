package main

import (
	"fmt"
)

/*
剑指 Offer 46. 把数字翻译成字符串

给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

提示：
- 0 <= num < 231

链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
*/
var testCases = []struct {
	nums   []int
	k      int
	result int
}{
	{[]int{2, 1}, 2, 1},
}

func main() {
	for _, tc := range testCases {
		got := 0
		want := tc.result
		if got != want {
			fmt.Printf("got: %d, want: %d\n", got, want)
		} else {
			fmt.Println("ok")
		}
	}

	//for _, tc := range testCases {
	//	got := searchRange(tc.nums, tc.target)
	//	want := tc.result
	//	ok := true
	//	for i := range want {
	//		if got[i] != want[i] {
	//			ok = false
	//			break
	//		}
	//	}
	//
	//	if ok {
	//		fmt.Println("ok")
	//	} else {
	//		fmt.Printf("got: %d, want: %d\n", got, want)
	//	}
	//}
}

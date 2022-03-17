package main

import (
	"fmt"
	v1 "github.com/youngzhu/go-leetcode/p005/v1"
)

/*
5. 最长回文子串

给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
- 1 <= s.length <= 1000
- s 仅由数字和英文字母组成

链接：https://leetcode-cn.com/problems/longest-palindromic-substring
*/
var testCases = []struct {
	s      string
	result string
}{
	{"cbbd", "bb"},
	{"bb", "bb"},
}

func main() {
	for _, tc := range testCases {
		got := v1.LongestPalindrome(tc.s)
		want := tc.result
		if got != want {
			fmt.Printf("got: %v, want: %v\n", got, want)
		} else {
			fmt.Println("ok")
		}
	}
}

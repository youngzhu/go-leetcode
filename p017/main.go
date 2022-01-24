package main

import (
	"fmt"

	. "github.com/youngzhu/go-leetcode/p017/v0"
)

/*
17. 电话号码的字母组合
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]

提示：
1. 0 <= digits.length <= 4
2. digits[i] 是范围 ['2', '9'] 的一个数字。

链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
*/
func main() {
	digits := "2"
	fmt.Println(digits, LetterCombinations(digits))

	digits = ""
	fmt.Println(digits, LetterCombinations(digits))

	digits = "23"
	fmt.Println(digits, LetterCombinations(digits))
}

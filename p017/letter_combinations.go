package main

import "fmt"

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
var digitLetters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)

	n := len(digits)
	if n == 0 {
		return result
	}

	letter := digitLetters[digits[0]]
	if n == 1 {
		result1 := make([]string, len(letter))
		for i, c := range letter {
			result1[i] = string(c)
		}
		return result1
	} else {
		for _, c := range letter {
			for _, s := range letterCombinations(digits[1:]) {
				result = append(result, string(c)+s)
			}
		}
		return result
	}
}

func main() {
	digits := "2"
	fmt.Println(digits, letterCombinations(digits))

	digits = ""
	fmt.Println(digits, letterCombinations(digits))

	digits = "23"
	fmt.Println(digits, letterCombinations(digits))
}

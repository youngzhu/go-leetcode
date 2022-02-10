package main

import (
	"fmt"
)

/*
22. 括号生成
数字n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8

链接：https://leetcode-cn.com/problems/generate-parentheses
*/
var ans []string

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	build("", n, n)
	return ans
}

// 参考别人的，自己没想法。。
func build(s string, left int, right int) {
	if left == 0 && right == 0 {
		ans = append(ans, s)
		return
	}
	if left == right {
		// 剩余的左括号和右括号相等，新的必须从 ( 开始
		build(s+"(", left-1, right)
	} else {
		// left < right
		// 剩余的左括号（left）不可能也不应该大于剩余的右括号（right）
		// 此时，后面可以跟 ( 也可以跟 )
		if left > 0 {
			build(s+"(", left-1, right)
		}
		build(s+")", left, right-1)
	}
}

func main() {
	n := 1
	fmt.Println(n, generateParenthesis(n))

	n = 3
	fmt.Println(n, generateParenthesis(n))
}

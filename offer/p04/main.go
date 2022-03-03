package main

import (
	"fmt"
	v2 "github.com/youngzhu/go-leetcode/offer/p04/v2"
)

/*
剑指 Offer 04. 二维数组中的查找

在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target=5，返回true。
给定 target=20，返回false。

限制：
- 0 <= n <= 1000
- 0 <= m <= 1000

链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
*/
var testCases = []struct {
	matrix [][]int
	target int
	result bool
}{
	//{[][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}, 5, true},
	//{[][]int{}, 0, false},
	//{[][]int{{}}, 1, false},
	//{[][]int{{-5}}, -5, true},
	//{[][]int{{1, 4}, {2, 5}}, 2, true},
	//{[][]int{{-1, 3}}, 1, false},
	//{[][]int{{-1, 3}}, 3, true},
	//{[][]int{{1, 2, 3, 4, 5}}, 3, true},
	//{[][]int{{1}, {2}, {3}, {4}, {5}}, 3, true},
	{[][]int{{2, 5}, {2, 8}, {7, 9}, {7, 11}, {9, 11}}, 7, true},
}

func main() {
	for _, tc := range testCases {
		got := v2.FindNumberIn2DArray(tc.matrix, tc.target)
		want := tc.result
		if got != want {
			fmt.Printf("got: %v, want: %v\n", got, want)
		} else {
			fmt.Println("ok")
		}
	}
}
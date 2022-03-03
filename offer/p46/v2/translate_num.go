package v1

import "strconv"

// 动态规划
// 计算dp[i]时最多只有dp[i-1]和dp[i-2]有用，其他的都不会再被使用
// 所以，可以将数组换成两个临时变量存储

// 奇怪的是，空间并没有变少。。。

func translateNum(num int) int {
	if num < 10 {
		return 1
	}

	s := strconv.Itoa(num)
	s = "0" + s
	n := len(s)

	// current dp[i]
	var c int
	// before, dp[i-1] dp[i-2]
	b1, b2 := 1, 1

	for i := 2; i < n; i++ {
		ss := string(s[i-1 : i+1])
		if ss >= "10" && ss <= "25" {
			c = b1 + b2
		} else {
			c = b1
		}
		// 更新b1,b2
		b2 = b1
		b1 = c
	}
	return c
}

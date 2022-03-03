package v1

import "strconv"

// 常规的动态规划（用数组）

func translateNum(num int) int {
	s := strconv.Itoa(num)
	s = "0" + s
	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n; i++ {
		ss := string(s[i-1 : i+1])
		if ss >= "10" && ss <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

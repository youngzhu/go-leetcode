package v2

/*
动态规划

本来是空间换时间的
结果跟v1比，空间大了，时间也更大了。。。
*/

type indexRange struct {
	begin, end int
}

func (r indexRange) length() int {
	return r.end - r.begin + 1
}

func LongestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// 初始化
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	var maxRange indexRange
	maxLen := 1
	var left, right int
	for right = 1; right < n; right++ {
		for left = 0; left < right; left++ {
			if s[left] != s[right] {
				dp[left][right] = false
			} else {
				if right-left <= 2 {
					dp[left][right] = true
				} else {
					dp[left][right] = dp[left+1][right-1]
				}
			}

			if dp[left][right] {
				theRange := indexRange{left, right}
				if theRange.length() > maxLen {
					maxRange = theRange
					maxLen = theRange.length()
				}
			}
		}
	}

	return s[maxRange.begin : maxRange.end+1]
}

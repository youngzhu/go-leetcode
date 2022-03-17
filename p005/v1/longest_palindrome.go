package v1

/*
中心扩散法

因为存在奇偶性，所以中心可能从1位或2位开始
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

	var theRange indexRange
	// 1位开始扩展
	for i := 1; i < n-1; i++ {
		rangeOne := extendAroundCenter(s, i, i)
		if rangeOne.length() > theRange.length() {
			theRange = rangeOne
		}
	}
	// 2位开始扩展
	for i := 0; i < n-1; i++ {
		rangeTwo := extendAroundCenter(s, i, i+1)
		if rangeTwo.length() > theRange.length() {
			theRange = rangeTwo
		}
	}

	return s[theRange.begin : theRange.end+1]
}

// 以[left, right]为中心扩散的回文起点和终点
func extendAroundCenter(s string, left, right int) indexRange {
	n := len(s)
	ll, rr := left, right
	for ll >= 0 && rr < n {
		if s[ll] == s[rr] {
			ll--
			rr++
		} else {
			break
		}
	}

	return indexRange{ll + 1, rr - 1}
}

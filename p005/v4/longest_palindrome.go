package v4

/*
Manacher + 中心扩散法
利用前者消除后者算法中的奇偶性
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

	t := manacher(s)
	nt := len(t)

	var theRange indexRange
	for i := 1; i < nt-1; i++ {
		rangeOne := extendAroundCenter(t, i, i)
		if rangeOne.length() > theRange.length() {
			theRange = rangeOne
		}
	}

	//begin := theRange.begin / 2
	//end := begin + theRange.length()/2
	//return s[begin:end]

	return s[theRange.begin/2 : (theRange.end)/2]
}

// 每个字符之间插入#
func manacher(s string) (t string) {
	n := len(s)

	t = "#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	t += "#"

	return
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

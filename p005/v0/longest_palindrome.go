package v0

import (
	"fmt"
	"sort"
)

/*
受到题解中的启发
可以将字符串倒转，然后求两个字符串的最长公共字子串

===
原来这个方法不靠谱啊，"aacabdkacaa"
刚还在想，我这么复制粘贴有啥意义
*/

func LongestPalindrome(s string) string {
	t := reverse(s)

	return longestCommonSubstring(s, t)
}

func longestCommonSubstring(s, t string) (lcs string) {
	suffix1 := NewSuffixArray(s)
	suffix2 := NewSuffixArray(t)

	// find the longest common substring by "merging" sorted suffixes
	lcs = ""
	i, j := 0, 0
	for i < suffix1.Length() && j < suffix2.Length() {
		p := suffix1.Index(i)
		q := suffix2.Index(j)
		x := lcpFrom(s, t, p, q)
		if len(x) > len(lcs) {
			lcs = x
		}
		if compare(s, t, p, q) < 0 {
			i++
		} else {
			j++
		}
	}

	return
}

type SuffixArray struct {
	suffixes []string
	n        int
}

func NewSuffixArray(txt string) SuffixArray {
	n := len(txt)
	suffixes := make([]string, n)
	for i := 0; i < n; i++ {
		suffixes[i] = txt[i:]
	}
	sort.Strings(suffixes)

	return SuffixArray{suffixes: suffixes, n: n}
}

// Length
// Returns the length of the input string
func (sa SuffixArray) Length() int {
	return sa.n
}

// Index
// Returns the index into the original string of the ith smallest suffix.
func (sa SuffixArray) Index(i int) int {
	if i < 0 || i >= sa.n {
		panic(`Illegal index: ` + fmt.Sprint(i))
	}
	return sa.n - len(sa.suffixes[i])
}

// LCP returns the length of the longest common prefix of the ith smallest suffix
// and the i-1 th smallest suffix.
func (sa SuffixArray) LCP(i int) int {
	if i < 1 || i >= sa.n {
		panic("Illegal argument")
	}
	return lcp(sa.suffixes[i], sa.suffixes[i-1])
}

// longest common prefix of s and t
func lcp(s, t string) int {
	n := min(len(s), len(t))
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			return i
		}
	}
	return n
}

// Select
// Returns the ith smallest suffix as a string
func (sa SuffixArray) Select(i int) string {
	if i < 0 || i >= sa.n {
		panic("Illegal argument")
	}
	return sa.suffixes[i]
}

// Rank
// Returns the number of suffixes strictly less than the query string.
// Note: rank(select(i)) == i
func (sa SuffixArray) Rank(query string) int {
	lo, hi := 0, sa.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midSuffix := sa.suffixes[mid]
		switch {
		case query < midSuffix:
			hi = mid - 1
		case query > midSuffix:
			lo = mid + 1
		default:
			return mid
		}
	}
	return lo
}

// return the longest common prefix of suffix s[p...] and suffix t[q...]
func lcpFrom(s, t string, p, q int) string {
	n := min(len(s)-p, len(t)-q)
	for i := 0; i < n; i++ {
		if s[p+i] != t[q+i] {
			return s[p : p+i]
		}
	}
	return s[p : p+n]
}

// compare suffix s[p...] and suffix t[q...]
func compare(s, t string, p, q int) int {
	pp, qq := len(s)-p, len(t)-q
	n := min(pp, qq)
	for i := 0; i < n; i++ {
		if s[p+i] != t[q+i] {
			return int(s[p+i]) - int(t[q+i])
		}
	}

	if pp < qq {
		return -1
	} else if pp > qq {
		return +1
	} else {
		return 0
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func reverse(s string) string {
	// 也可以用rune，只有字母和数字byte够用了
	bytes := []byte(s)
	n := len(bytes)
	for i, j := 0, n-1; i < j; {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}

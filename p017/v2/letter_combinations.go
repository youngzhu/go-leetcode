package v2

// 回溯
// 常量改成字符串数组看看有没有区别
// 结果：没啥区别

var digitLetters = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

var result []string

func LetterCombinations(digits string) []string {
	result = make([]string, 0)

	if digits == "" {
		return result
	}

	backtrack("", digits)

	return result
}

func backtrack(combination, digits string) {
	if digits == "" {
		result = append(result, combination)
	} else {
		d := digits[0]
		for _, c := range digitLetters[d] {
			backtrack(combination+c, digits[1:])
		}
	}
}

package v1

// 回溯

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
			backtrack(combination+string(c), digits[1:])
		}
	}
}

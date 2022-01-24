package v0

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

func LetterCombinations(digits string) []string {
	result := make([]string, 0)

	n := len(digits)
	if n == 0 {
		return result
	}

	letter := digitLetters[digits[0]]
	if n == 1 {
		result1 := make([]string, len(letter))
		for i, c := range letter {
			result1[i] = string(c)
		}
		return result1
	} else {
		for _, c := range letter {
			for _, s := range LetterCombinations(digits[1:]) {
				result = append(result, string(c)+s)
			}
		}
		return result
	}
}

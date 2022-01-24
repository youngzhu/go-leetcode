package queue

import "github.com/youngzhu/algs4-go/fund"

/*
队列
应该可行。但LeetCode不能借助外力
*/
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
	queue := fund.NewQueue()
	queue.Enqueue("")

	for _, c := range digits {
		letters := digitLetters[byte(c)]
		n := queue.Size()
		for i := 0; i < n; i++ {
			s0 := queue.Dequeue().(string)
			for i := 0; i < len(letters); i++ {
				queue.Enqueue(s0 + string(letters[i]))
			}
		}
	}

	result := make([]string, queue.Size())
	for i, s := range queue.Iterator() {
		result[i] = s.(string)
	}
	return result
}

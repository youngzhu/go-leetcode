package v3

import "sort"

// 在v2的基础上，改用二维数组
// 结果：效率更低，时间多了，空间也多了

const max = 10001 // 假设一个最大值（不可能达到，因为amount最大为10000）

var (
	ans   int
	cache [][]int // 缓存数据，避免重复的和无效的计算
)

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	ans = max
	cache = make([][]int, amount+1)
	for i := range cache {
		s := make([]int, len(coins))
		for j := range s {
			s[j] = max
		}
		cache[i] = s
	}

	minCount(coins, amount, len(coins)-1, 0)

	if ans != max {
		return ans
	}
	return -1
}

func minCount(coins []int, amount int, i int, count int) {
	if amount == 0 {
		ans = min(ans, count)
		return
	}
	if i < 0 {
		return
	}
	if cache[amount][i] <= count {
		return
	}

	for j := amount / coins[i]; j >= 0 && j+count < ans; j-- {
		amtLeft := amount - j*coins[i]
		minCount(coins, amtLeft, i-1, count+j)
	}
	if count < cache[amount][i] {
		cache[amount][i] = count
	}
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

package v1

import "sort"

const max = 10001 // 假设一个最大值（不可能达到，因为amount最大为10000）

var (
	ans   int
	cache map[int]int // 缓存数据，避免重复的和无效的计算
)

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	ans = max
	cache = make(map[int]int)

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
	val, ok := cache[amount]
	if ok && val <= count {
		return
	}

	for j := amount / coins[i]; j >= 0 && j+count < ans; j-- {
		amtLeft := amount - j*coins[i]
		minCount(coins, amtLeft, i-1, count+j)
	}
	val, ok = cache[amount]
	if ok {
		if count < val {
			cache[amount] = count
		}
	} else {
		cache[amount] = count
	}
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

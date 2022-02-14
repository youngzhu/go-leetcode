package v5

import "sort"

// 在 v4 基础上优化
// 1 将硬币排序
// 2 减少部分循环，例如 coins[i]已经超出了，就没必要看之后的了

// 结果：几乎没啥影响

const max = 10001 // 假设一个最大值（不可能达到，因为amount最大为10000）

var (
	cache []int // 缓存数据，避免重复的和无效的计算
)

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 初始化
	cache = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		cache[i] = max
	}
	cache[0] = 0

	sort.Ints(coins)

	// 动态规划
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// 总金额 i 分为两个部分：当前硬币所表示的面额，和剩余的金额
			// 当前硬币 固定是1个
			// 所以，相当于找出 cache[left] 的最小值
			left := i - coin
			if left < 0 {
				break
			}

			cache[i] = min(cache[i], cache[left]+1)
		}
	}

	// 结束
	if cache[amount] == max {
		return -1
	}
	return cache[amount]
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

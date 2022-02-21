package v2

// CanJump
// 记录并更新每一步能够到达的最远位置
// 如果longest>=n-1，则可以到达终点
func CanJump(nums []int) bool {
	n := len(nums)

	longest := nums[0]
	// 本来猜想 range 是否更高效。
	// 结果：并没有
	//for i := 0; i < n; i++ {
	//	if i <= longest { // 说明i可到达
	//		longest = max(longest, i+nums[i])
	//	}
	//}
	for i, num := range nums {
		if i <= longest { // 说明i可到达
			longest = max(longest, i+num)
		}
		if longest >= n-1 {
			return true
		}
	}

	return false
}

func max(i, j int) int {
	if j > i {
		return j
	}
	return i
}

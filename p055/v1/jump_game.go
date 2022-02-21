package v1

// CanJump 检查是否有0元素，如果没有，则一定能到达终点
// 如果有，则看能否跳过这个元素
func CanJump(nums []int) bool {
	n := len(nums)

	if n == 1 {
		return true
	}

	// 找0，最后一个是否为0不重要
	endIdx := n - 1
	i := 0
	for i < endIdx {
		if nums[i] == 0 {
			if !canReach(nums, i-1, i+1) {
				return false
			}
		}
		i++
	}

	return true
}

// 从[0, end]的范围内，能否到达target
func canReach(nums []int, end int, target int) bool {
	for i := end; i >= 0; i-- {
		if nums[i] >= target-i {
			return true
		}
	}

	return false
}

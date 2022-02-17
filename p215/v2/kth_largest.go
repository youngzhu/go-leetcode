package v2

// 在上一版本的基础上，不复制数组（root从0开始），看看是否能节省空间
// 结果：稍有进步， 3.3m vs 3.5m

func FindKthLargest(nums []int, k int) int {
	n := len(nums)

	// 生成堆
	for parent := (n - 1) / 2; parent >= 0; parent-- {
		sink(nums, parent, n-1)
	}

	i := n - 1
	for k > 0 {
		nums[0], nums[i] = nums[i], nums[0]
		i--
		sink(nums, 0, i)

		k--
	}

	return nums[i+1]
}

// 下沉，使得
// nums[parent]>=nums[leftChild]
// nums[parent]>=nums[rightChild]
func sink(nums []int, parent, n int) {
	for {
		leftChild := 2*parent + 1
		if leftChild > n {
			break
		}
		rightChild := leftChild + 1
		elderChild := leftChild // left right 中更大的那个
		if rightChild <= n && nums[rightChild] > nums[leftChild] {
			elderChild = rightChild
		}
		if nums[parent] >= nums[elderChild] {
			break
		}
		// child更大，则交换
		nums[parent], nums[elderChild] = nums[elderChild], nums[parent]
		parent = elderChild
	}
}

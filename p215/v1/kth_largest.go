package v1

func FindKthLargest(nums []int, k int) int {
	n := len(nums)

	// 下标从1开始，方便计算
	heap := make([]int, n+1)
	copy(heap[1:n+1], nums)

	// 生成堆
	for parent := n / 2; parent >= 1; parent-- {
		sink(heap, parent, n)
	}

	i := n
	for k > 0 {
		heap[1], heap[i] = heap[i], heap[1]
		i--
		sink(heap, 1, i)

		k--
	}

	return heap[i+1]
}

// 下沉，使得
// nums[parent]>=nums[leftChild]
// nums[parent]>=nums[rightChild]
func sink(nums []int, parent, n int) {
	for 2*parent <= n {
		leftChild := 2 * parent
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

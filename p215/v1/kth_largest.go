package v1

// 堆排序 based 1
/*
### 解题思路
堆排序的特点：不用完全排序就能找到最大元素
题目的要求：找出前K个最大数（其他可以是无序的）
两者非常契合

### 二叉堆
堆的两种数组表示
#### base 1（root=arr[1]）
**父节点 p**
leftChild=p\*2
rightChild=p\*2+1
**任意节点 i**
parent=i/2

优点：容易理解
缺点：需要额外的空间

#### base 0（root=arr[0]）
**父节点 p**
leftChild=p\*2+1
rightChild=p\*2+2
**任意节点 i**
parent=(i-1)/2

优点：原地排序，不需要额外空间
缺点：稍微有点绕（边界值问题）
*/

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

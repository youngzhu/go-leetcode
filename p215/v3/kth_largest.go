package v3

// 堆实现
// 结果：同v2效果相当
/*
### 解题思路
1. 设计并实现一个最小堆，堆的最大容量为k
2. 先用前k个元素将堆填满
3. k之后的元素，与堆顶元素比较。如果num[i]>heap.top，则从heap中删除一个元素，num[i]入堆

关键点：前K个最大的元素中，最小的那一个就是题目要求的第K大的元素，即堆的顶元素
*/

func FindKthLargest(nums []int, k int) int {
	heap := newMinimumHeap(k)

	for _, num := range nums {
		if k > 0 { // 先将堆填满
			heap.Insert(num)
			k--
		} else {
			// 用更大的值将堆中最小的值替换出去
			if num > heap.GetMinimum() {
				heap.Remove()
				heap.Insert(num)
			}
		}
	}

	// 前K个最大的元素中，最小的那一个
	// 即是题目要求的第K大的元素
	return heap.GetMinimum()
}

type MinimumHeap struct {
	items    []int
	capacity int // 容量
	size     int // 实际大小
}

func newMinimumHeap(n int) *MinimumHeap {
	items := make([]int, n)
	return &MinimumHeap{items: items, capacity: n, size: 0}
}

func (h *MinimumHeap) Insert(item int) {
	if h.IsFull() {
		panic("heap is full")
	}

	h.items[h.size] = item
	h.swim(h.size)
	h.size++
}

func (h *MinimumHeap) swim(i int) {
	parent := h.getParent(i)

	if parent != -1 && h.GetItem(i) < h.GetItem(parent) {
		h.swap(parent, i)
		h.swim(parent)
	}
}

func (h *MinimumHeap) Remove() int {
	root := h.items[0]

	h.items[0] = h.GetItem(h.size - 1)
	h.size--
	h.sink(0)

	return root
}

func (h *MinimumHeap) sink(i int) {
	leftChild := h.getLeftChild(i)
	rightChild := h.getRightChild(i)

	// 没有左节点，更不会有右节点。所以是叶子节点
	if leftChild != -1 {
		smaller := leftChild
		if rightChild != -1 && h.GetItem(rightChild) < h.GetItem(leftChild) {
			smaller = rightChild
		}

		if h.GetItem(smaller) < h.GetItem(i) {
			h.swap(i, smaller)
			h.sink(smaller)
		}
	}
}

func (h *MinimumHeap) getLeftChild(p int) int {
	left := 2*p + 1
	if left >= h.size {
		return -1
	}

	return left
}

func (h *MinimumHeap) getRightChild(p int) int {
	right := 2*p + 2
	if right >= h.size {
		return -1
	}

	return right
}

func (h *MinimumHeap) getParent(i int) int {
	if i <= 0 || i >= h.capacity {
		return -1
	}
	return (i - 1) / 2
}

func (h *MinimumHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MinimumHeap) GetMinimum() int {
	if h.IsEmpty() {
		panic("heap is empty")
	}
	return h.items[0]
}

func (h *MinimumHeap) GetItem(index int) int {
	if index < 0 || index >= h.capacity {
		panic("out of range")
	}
	return h.items[index]
}

func (h *MinimumHeap) IsFull() bool {
	return h.size == h.capacity
}

func (h *MinimumHeap) IsEmpty() bool {
	return h.size == 0
}

package v3

// 堆实现
// 结果：稍有进步， 3.3m vs 3.5m

func FindKthLargest(nums []int, k int) int {
	heap := newMinimumHeap(k)

	for _, num := range nums {
		if k > 0 {
			heap.Insert(num)
			k--
		} else {
			if num > heap.GetMinimum() {
				heap.Remove()
				heap.Insert(num)
			}
		}
	}

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
	if i <= 0 || i >= h.size {
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

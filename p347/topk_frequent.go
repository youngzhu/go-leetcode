package main

import "fmt"

/*
347. 前 K 个高频元素

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]

提示：
- 1 <= nums.length <= 10^5
- k 的取值范围是 [1, 数组中不相同的元素的个数]
- 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n是数组大小。

链接：https://leetcode-cn.com/problems/top-k-frequent-elements
*/
func topKFrequent(nums []int, k int) []int {
	// 使用map计算频率
	// key: num; value: frequent
	numFrequentMap := make(map[int]int)

	var frequent int
	for _, num := range nums {
		v, ok := numFrequentMap[num]
		if ok {
			frequent = v + 1
		} else {
			frequent = 1
		}
		numFrequentMap[num] = frequent
	}
	// 只有一个数字
	if len(numFrequentMap) == 1 {
		return nums[0:1]
	}

	// key: frequent 防止不同数字又相同的频率，保证频率的唯一性
	frequentMap := make(map[int]bool)
	for _, v := range numFrequentMap {
		frequentMap[v] = true
	}

	//fmt.Println("numFrequentMap:", numFrequentMap)
	//fmt.Println("frequentMap:", frequentMap)

	// 把频率放入数组中，找出前k个最大频率
	// 这样，就跟 215. 数组中的第K个最大元素 类似
	arr := make([]int, len(frequentMap))
	i := 0
	for key, _ := range frequentMap {
		arr[i] = key
		i++
	}
	topK := findTopK(arr, k)

	// 反向查找，根据频率找到数字
	ans := make([]int, k)
	count := 0
	// 频率一定要从大到小
	// 可能不同的数字具有相同的频率
	for _, f := range topK {
		for num, v := range numFrequentMap {
			if count == k {
				break
			}
			if v == f {
				ans[count] = num
				count++
			}
		}
	}

	return ans
}

func findTopK(nums []int, k int) []int {
	n := len(nums)

	// 生成堆
	for parent := (n - 1) / 2; parent >= 0; parent-- {
		sink(nums, parent, n-1)
	}

	topK := make([]int, k)
	count := 0
	i := n - 1
	for count < k && i >= 0 {
		topK[count] = nums[0]
		count++
		nums[0], nums[i] = nums[i], nums[0]
		i--
		sink(nums, 0, i)
	}

	//fmt.Println("topK:", topK)
	return topK
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

func main() {
	var ans []int

	//ans = topKFrequent([]int{1, 2}, 2)
	//fmt.Println(ans)
	//
	//ans = topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)
	//fmt.Println(ans) // [-1,2]

	ans = topKFrequent([]int{5, 3, 1, 1, 1, 3, 73, 1}, 2)
	fmt.Println(ans) // [1,3]
}

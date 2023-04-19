package main

import "container/heap"

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
// 
//
// 示例 2: 
//
// 
//输入: nums = [1], k = 1
//输出: [1] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// k 的取值范围是 [1, 数组中不相同的元素的个数] 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的 
// 
//
// 
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。 
//
// 👍 1536 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type myHeap [][2]int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i][1] < (*h)[j][1]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.([2]int))
}

func topKFrequent(nums []int, k int) []int {
	cntMap := map[int]int{}
	for _, num := range nums {
		cntMap[num] ++
	}
	h := new(myHeap)
	for num, cnt := range cntMap {
		heap.Push(h, [2]int{num, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := k-1; i >= 0; i-- {
		result[i] = heap.Pop(h).([2]int)[0]
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

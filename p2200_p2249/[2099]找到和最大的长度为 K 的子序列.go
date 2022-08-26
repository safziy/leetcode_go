package main

import (
	"container/heap"
)

//给你一个整数数组 nums 和一个整数 k 。你需要找到 nums 中长度为 k 的 子序列 ，且这个子序列的 和最大 。
//
// 请你返回 任意 一个长度为 k 的整数子序列。 
//
// 子序列 定义为从一个数组里删除一些元素后，不改变剩下元素的顺序得到的数组。 
//
// 
//
// 示例 1： 
//
// 输入：nums = [2,1,3,3], k = 2
//输出：[3,3]
//解释：
//子序列有最大和：3 + 3 = 6 。 
//
// 示例 2： 
//
// 输入：nums = [-1,-2,3,4], k = 3
//输出：[-1,3,4]
//解释：
//子序列有最大和：-1 + 3 + 4 = 6 。
// 
//
// 示例 3： 
//
// 输入：nums = [3,4,3,3], k = 2
//输出：[3,4]
//解释：
//子序列有最大和：3 + 4 = 7 。
//另一个可行的子序列为 [4, 3] 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 1000 
// -10⁵ <= nums[i] <= 10⁵ 
// 1 <= k <= nums.length 
// 
//
// Related Topics 数组 哈希表 排序 堆（优先队列） 👍 22 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
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
	*h = append(*h, v.(int))
}

func maxSubsequence(nums []int, k int) []int {
	h := new(myHeap)
	l := len(nums)
	for i := 0; i < l; i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	min, cnt := heap.Pop(h).(int), 1
	for h.Len() > 0 && heap.Pop(h).(int) == min {
		cnt ++
	}
	result := make([]int, 0, k)
	for i := 0; i < l; i++ {
		if nums[i] > min {
			result = append(result, nums[i])
		} else if nums[i] == min && cnt > 0 {
			result = append(result, nums[i])
			cnt --
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

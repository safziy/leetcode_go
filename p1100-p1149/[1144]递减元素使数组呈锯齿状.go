package main

import "math"

//给你一个整数数组 nums，每次 操作 会从中选择一个元素并 将该元素的值减少 1。
//
// 如果符合下列情况之一，则数组 A 就是 锯齿数组：
//
//
// 每个偶数索引对应的元素都大于相邻的元素，即 A[0] > A[1] < A[2] > A[3] < A[4] > ...
// 或者，每个奇数索引对应的元素都大于相邻的元素，即 A[0] < A[1] > A[2] < A[3] > A[4] < ...
//
//
// 返回将数组 nums 转换为锯齿数组所需的最小操作次数。
//
//
//
// 示例 1：
//
// 输入：nums = [1,2,3]
//输出：2
//解释：我们可以把 2 递减到 0，或把 3 递减到 1。
//
//
// 示例 2：
//
// 输入：nums = [9,6,1,6,2]
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 1000
//
//
// 👍 73 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	get := func(idx int) int {
		if idx >= 0 && idx < n {
			return nums[idx]
		}
		return math.MaxInt32
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// res1 表示 a[0] > a[1] < a[2] > a[3]的解,res2 表示 a[0] < a[1] > a[2] < a[3]的解
	res1, res2 := 0, 0
	for i := 0; i < n; i++ {
		nestMin := min(get(i-1), get(i+1))
		if i %2 == 0 {
			res1 += nums[i] - min(nums[i], nestMin-1)
		} else {
			res2 += nums[i] - min(nums[i], nestMin-1)
		}
	}
	return min(res1, res2)
}

//leetcode submit region end(Prohibit modification and deletion)

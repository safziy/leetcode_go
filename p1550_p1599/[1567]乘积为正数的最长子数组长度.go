package main

//给你一个整数数组 nums ，请你求出乘积为正数的最长子数组的长度。
//
// 一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。
//
// 请你返回乘积为正数的最长子数组长度。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-2,-3,4]
//输出：4
//解释：数组本身乘积就是正数，值为 24 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,-2,-3,-4]
//输出：3
//解释：最长乘积为正数的子数组为 [1,-2,-3] ，乘积为 6 。
//注意，我们不能把 0 也包括到子数组中，因为这样乘积为 0 ，不是正数。
//
// 示例 3：
//
//
//输入：nums = [-1,-2,-3,0,1]
//输出：2
//解释：乘积为正数的最长子数组是 [-1,-2] 或者 [-2,-3] 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
//
//
//
//
// 👍 205 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func getMaxLen(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// positive 当前结果为正时的长度， negative即结果为负的时候的长度
	positive, negative, n := 0, 0, len(nums)
	if nums[0] > 0 {
		positive++
	} else if nums[0] < 0 {
		negative++
	}
	result := positive
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			if negative > 0 {
				positive, negative = positive+1, negative+1
			} else {
				positive, negative = positive+1, 0
			}
		} else if nums[i] < 0 {
			if negative > 0 {
				positive, negative = negative+1, positive+1
			} else {
				positive, negative = 0, positive+1
			}
		} else {
			positive, negative = 0, 0
		}
		result = max(result, positive)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

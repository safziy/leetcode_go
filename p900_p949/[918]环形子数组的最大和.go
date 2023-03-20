package main

import "math"

//给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
//
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个
//元素是 nums[(i - 1 + n) % n] 。
//
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不
//存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-2,3,-2]
//输出：3
//解释：从子数组 [3] 得到最大和 3
//
//
// 示例 2：
//
//
//输入：nums = [5,-3,5]
//输出：10
//解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
//
//
// 示例 3：
//
//
//输入：nums = [3,-2,2,-3]
//输出：3
//解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
//
//
// 👍 466 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubarraySumCircular(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	kadane := func(s, e int, sign int) int {
		ans, cur := math.MinInt32, math.MinInt32
		for i := s; i <= e; i++ {
			cur = sign*nums[i] + max(cur, 0)
			ans = max(ans, cur)
		}
		return ans
	}
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}
	ans1 := kadane(0, n-1, 1)
	ans2 := sum + kadane(1, n-1, -1)
	ans3 := sum + kadane(0, n-2, -1)
	return max(ans1, max(ans2, ans3))
}

//leetcode submit region end(Prohibit modification and deletion)

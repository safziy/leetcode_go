package main

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// 👍 3023 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	// 方法一：动态规划 时间复杂度O(n²) 空间复杂度O(n)
	//n, result := len(nums), 1
	//// dp[i] 表示以字符i为结束字符的最大序列
	//dp := make([]int, n)
	//for i := range dp {
	//	dp[i] = 1
	//}
	//for i := 1; i < n; i++ {
	//	for j := 0; j < i; j++ {
	//		if nums[i] > nums[j] && dp[j] >= dp[i] {
	//			dp[i] = dp[j] + 1
	//		}
	//	}
	//	if dp[i] > result {
	//		result = dp[i]
	//	}
	//}
	//return result

	// 方法二：贪心+二分查找 时间复杂度O(n㏒n) 空间复杂度 O(n)
	n, result := len(nums), 1
	// d[i] 表示长度为i的最大序列的最后一个值
	d := make([]int, n+1)
	d[result] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[result] {
			result++
			d[result] = nums[i]
		} else {
			l, r, pos := 1, result, 0
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

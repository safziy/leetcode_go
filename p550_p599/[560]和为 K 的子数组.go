package main
//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,1], k = 2
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3], k = 3
//输出：2
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2 * 10⁴ 
// -1000 <= nums[i] <= 1000 
// -10⁷ <= k <= 10⁷ 
// 
//
// 👍 1877 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	// 方法一: 前缀和+hashMap
	result, preSum, sumMap := 0, 0, map[int]int{}
	sumMap[0] = 1
	for _, num := range nums {
		preSum += num
		if cnt, ok := sumMap[preSum-k]; ok {
			result += cnt
		}
		sumMap[preSum] += 1
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

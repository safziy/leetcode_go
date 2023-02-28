package main

//给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
//
// 由于答案可能很大，因此 返回答案模 10^9 + 7 。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,1,2,4]
//输出：17
//解释：
//子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
//最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
//
// 示例 2：
//
//
//输入：arr = [11,81,94,43,3]
//输出：444
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 3 * 10⁴
// 1 <= arr[i] <= 3 * 10⁴
//
//
//
//
// Related Topics 栈 数组 动态规划 单调栈 👍 534 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func sumSubarrayMins(arr []int) int {
	const mod int = 1e9 + 7
	// stack维护单调递增栈， dp[i]为arr[i] 为最右且最小的最长子序列长度为k
	result, stack, dp := 0, []int{}, make([]int, len(arr))
	for i, num := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > num {
			stack = stack[:len(stack)-1]
		}
		k := i + 1
		if len(stack) > 0 {
			k = i - stack[len(stack)-1]
		}
		dp[i] = k * num
		if len(stack) > 0 {
			dp[i] += dp[i-k]
		}
		result = (result + dp[i]) % mod
		stack = append(stack, i)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

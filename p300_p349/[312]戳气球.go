package main

//有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i -
// 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。 
//
// 求所能获得硬币的最大数量。 
//
// 
//示例 1：
//
// 
//输入：nums = [3,1,5,8]
//输出：167
//解释：
//nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
//coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167 
//
// 示例 2： 
//
// 
//输入：nums = [1,5]
//输出：10
// 
//
// 
//
// 提示： 
//
// 
// n == nums.length 
// 1 <= n <= 300 
// 0 <= nums[i] <= 100 
// 
//
// Related Topics 数组 动态规划 👍 1086 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	val := func(idx int) int {
		if idx < 0 || idx >= n {
			return 1
		}
		return nums[idx]
	}
	for s := 2; s < n+2; s++ {
		for i := 0; i < n+2-s; i++ {
			for k := i+1; k < i+s; k++ {
				temp := dp[i][k] + val(i-1) * val(k-1) * val(i+s-1) + dp[k][i+s]
				if temp > dp[i][i+s] {
					dp[i][i+s] = temp
				}
			}
		}
	}
	return dp[0][n+1]
}
//leetcode submit region end(Prohibit modification and deletion)

package main

//给你一个长度为 n 的整数数组 nums ，这个数组中至多有 50 个不同的值。同时你有 m 个顾客的订单 quantity ，其中，整数
//quantity[i] 是第 i 位顾客订单的数目。请你判断是否能将 nums 中的整数分配给这些顾客，且满足：
//
//
// 第 i 位顾客 恰好 有 quantity[i] 个整数。
// 第 i 位顾客拿到的整数都是 相同的 。
// 每位顾客都满足上述两个要求。
//
//
// 如果你可以分配 nums 中的整数满足上面的要求，那么请返回 true ，否则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4], quantity = [2]
//输出：false
//解释：第 0 位顾客没办法得到两个相同的整数。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,3], quantity = [2]
//输出：true
//解释：第 0 位顾客得到 [3,3] 。整数 [1,2] 都没有被使用。
//
//
// 示例 3：
//
//
//输入：nums = [1,1,2,2], quantity = [2,2]
//输出：true
//解释：第 0 位顾客得到 [1,1] ，第 1 位顾客得到 [2,2] 。
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= 1000
// m == quantity.length
// 1 <= m <= 10
// 1 <= quantity[i] <= 10⁵
// nums 中至多有 50 个不同的数字。
//
//
// Related Topics 位运算 数组 动态规划 回溯 状态压缩 👍 44 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canDistribute(nums []int, quantity []int) bool {
	m := 1 << len(quantity)
	// 计算第j(状态子集)个客户需要的总数量
	sum := make([]int, m)
	for i, q := range quantity {
		iBit := 1 << i
		for j := 0; j < iBit; j++ {
			sum[iBit|j] = sum[j] + q
		}
	}
	// 统计一下
	cntMap := map[int]int{}
	for _, num := range nums {
		cntMap[num]++
	}
	n := len(cntMap)
	// dp[i][j] 代表i个元素是否能满足集合为j的客户
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m)
		dp[i][0] = true
	}
	i := 0
	for _, cnt := range cntMap {
		for j := 0; j < m; j++ {
			if dp[i][j] {
				dp[i+1][j] = true
				continue
			}
			for k := j; k > 0; k = (k - 1) & j {
				if sum[k] <= cnt && dp[i][j^k] {
					dp[i+1][j] = true
					break
				}
			}
		}
		i++
	}
	return dp[n][m-1]
}

//leetcode submit region end(Prohibit modification and deletion)

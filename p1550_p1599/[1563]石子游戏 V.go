package main

//几块石子 排成一行 ，每块石子都有一个关联值，关联值为整数，由数组 stoneValue 给出。
//
// 游戏中的每一轮：Alice 会将这行石子分成两个 非空行（即，左侧行和右侧行）；Bob 负责计算每一行的值，即此行中所有石子的值的总和。Bob 会丢弃值最
//大的行，Alice 的得分为剩下那行的值（每轮累加）。如果两行的值相等，Bob 让 Alice 决定丢弃哪一行。下一轮从剩下的那一行开始。 
//
// 只 剩下一块石子 时，游戏结束。Alice 的分数最初为 0 。 
//
// 返回 Alice 能够获得的最大分数 。 
//
// 
//
// 示例 1： 
//
// 输入：stoneValue = [6,2,3,4,5,5]
//输出：18
//解释：在第一轮中，Alice 将行划分为 [6，2，3]，[4，5，5] 。左行的值是 11 ，右行的值是 14 。Bob 丢弃了右行，Alice 的分数现
//在是 11 。
//在第二轮中，Alice 将行分成 [6]，[2，3] 。这一次 Bob 扔掉了左行，Alice 的分数变成了 16（11 + 5）。
//最后一轮 Alice 只能将行分成 [2]，[3] 。Bob 扔掉右行，Alice 的分数现在是 18（16 + 2）。游戏结束，因为这行只剩下一块石头了。
//
// 
//
// 示例 2： 
//
// 输入：stoneValue = [7,7,7,7,7,7,7]
//输出：28
// 
//
// 示例 3： 
//
// 输入：stoneValue = [4]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= stoneValue.length <= 500 
// 1 <= stoneValue[i] <= 10^6 
// 
//
// 👍 49 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func stoneGameV(stoneValue []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(stoneValue)
	dp := make([][]int, n)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + stoneValue[i-1]
		dp[i-1] = make([]int, n)
	}
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left == right {
			return 0
		}
		if dp[left][right] != 0 {
			return dp[left][right]
		}
		for i := left; i < right; i++ {
			leftSum, rightSum := prefixSum[i+1]-prefixSum[left], prefixSum[right+1]-prefixSum[i+1]
			if leftSum > rightSum {
				dp[left][right] = max(dp[left][right], dfs(i+1, right)+rightSum)
			} else if leftSum < rightSum {
				dp[left][right] = max(dp[left][right], dfs(left, i)+leftSum)
			} else {
				dp[left][right] = max(dp[left][right], max(dfs(i+1, right), dfs(left, i))+rightSum)
			}
		}
		return dp[left][right]
	}
	return dfs(0, n-1)
}

//leetcode submit region end(Prohibit modification and deletion)

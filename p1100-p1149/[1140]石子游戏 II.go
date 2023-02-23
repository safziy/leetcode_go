package main

//爱丽丝和鲍勃继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
//
// 爱丽丝和鲍勃轮流进行，爱丽丝先开始。最初，M = 1。
//
// 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
//
// 游戏一直持续到所有石子都被拿走。
//
// 假设爱丽丝和鲍勃都发挥出最佳水平，返回爱丽丝可以得到的最大数量的石头。
//
//
//
// 示例 1：
//
//
//输入：piles = [2,7,9,4,4]
//输出：10
//解释：如果一开始Alice取了一堆，Bob取了两堆，然后Alice再取两堆。爱丽丝可以得到2 + 4 + 4 = 10堆。如果Alice一开始拿走了两堆，那
//么Bob可以拿走剩下的三堆。在这种情况下，Alice得到2 + 7 = 9堆。返回10，因为它更大。
//
//
// 示例 2:
//
//
//输入：piles = [1,2,3,4,5,100]
//输出：104
//
//
//
//
// 提示：
//
//
// 1 <= piles.length <= 100
//
// 1 <= piles[i] <= 10⁴
//
//
// 👍 179 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func stoneGameII(piles []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(piles)
	// dp[i][j]表示剩余[i : len - 1]堆时，M = j的情况下，先取的人能获得的最多石子数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}
	suffixSum := 0
	for i := n - 1; i >= 0; i-- {
		suffixSum += piles[i]
		for m := 1; m <= n; m++ {
			if 2*m >= n-i {
				dp[i][m] = suffixSum
			} else {
				for x := 1; x <= 2*m; x++ {
					dp[i][m] = max(dp[i][m], suffixSum-dp[i+x][max(m, x)])
				}
			}
		}
	}
	return dp[0][1]
}

//leetcode submit region end(Prohibit modification and deletion)

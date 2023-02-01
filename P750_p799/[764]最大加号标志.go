package main

//在一个 n x n 的矩阵 grid 中，除了在数组 mines 中给出的元素为 0，其他每个元素都为 1。mines[i] = [xi, yi]表示
//grid[xi][yi] == 0
//
// 返回 grid 中包含 1 的最大的 轴对齐 加号标志的阶数 。如果未找到加号标志，则返回 0 。
//
// 一个 k 阶由 1 组成的 “轴对称”加号标志 具有中心网格 grid[r][c] == 1 ，以及4个从中心向上、向下、向左、向右延伸，长度为 k-1，
//由 1 组成的臂。注意，只有加号标志的所有网格要求为 1 ，别的网格可能为 0 也可能为 1 。
//
//
//
// 示例 1：
//
//
//
//
//输入: n = 5, mines = [[4, 2]]
//输出: 2
//解释: 在上面的网格中，最大加号标志的阶只能是2。一个标志已在图中标出。
//
//
// 示例 2：
//
//
//
//
//输入: n = 1, mines = [[0, 0]]
//输出: 0
//解释: 没有加号标志，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 500
// 1 <= mines.length <= 5000
// 0 <= xi, yi < n
// 每一对 (xi, yi) 都 不重复
//
//
// Related Topics 数组 动态规划 👍 159 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func orderOfLargestPlusSign(n int, mines [][]int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = n
		}
	}
	for _, mine := range mines {
		dp[mine[0]][mine[1]] = 0
	}
	count := func(row, col, cnt int) int {
		if dp[row][col] == 0 {
			return 0
		} else {
			return cnt + 1
		}
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			cnt = count(i, j, cnt)
			dp[i][j] = min(dp[i][j], cnt)
		}
		cnt = 0
		for j := n - 1; j >= 0; j-- {
			cnt = count(i, j, cnt)
			dp[i][j] = min(dp[i][j], cnt)
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			cnt = count(j, i, cnt)
			dp[j][i] = min(dp[j][i], cnt)
		}
		cnt = 0
		for j := n - 1; j >= 0; j-- {
			cnt = count(j, i, cnt)
			dp[j][i] = min(dp[j][i], cnt)
			ans = max(dp[j][i], ans)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

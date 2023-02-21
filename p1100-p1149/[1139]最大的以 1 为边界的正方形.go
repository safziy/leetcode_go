package main

//给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回 0
//。
//
//
//
// 示例 1：
//
// 输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
//输出：9
//
//
// 示例 2：
//
// 输入：grid = [[1,1,0,0]]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= grid.length <= 100
// 1 <= grid[0].length <= 100
// grid[i][j] 为 0 或 1
//
//
// 👍 136 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func largest1BorderedSquare(grid [][]int) int {
	min := func(a, b int) int {
		if a < b { return a}
		return b
	}
	m, n := len(grid), len(grid[0])
	up, left := make([][]int, m+1), make([][]int, m+1)
	for i := 0; i <= m; i++ {
		up[i] = make([]int, n+1)
		left[i] = make([]int, n+1)
	}
	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				up[i][j] = up[i-1][j] + 1
				left[i][j] = left[i][j-1] + 1
				border := min(up[i][j], left[i][j])
				for up[i][j+1-border] < border || left[i+1-border][j] < border {
					border --
				}
				if border > result {
					result = border
				}
			}
		}
	}
	return result * result
}

//leetcode submit region end(Prohibit modification and deletion)

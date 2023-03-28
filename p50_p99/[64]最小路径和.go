package main

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。 
//
// 
//
// 示例 1： 
// 
// 
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
// 
//
// 示例 2： 
//
// 
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 200 
// 0 <= grid[i][j] <= 100 
// 
//
// 👍 1460 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	m, n := len(grid), len(grid[0])
	for r := 0; r < m; r++ {
		if r > 0 {
			grid[r][0] += grid[r-1][0]
		}
		for c := 1; c < n; c++ {
			if r > 0 {
				grid[r][c] += min(grid[r-1][c], grid[r][c-1])
			} else {
				grid[r][c] += grid[r][c-1]
			}
		}
	}
	return grid[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

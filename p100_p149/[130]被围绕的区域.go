package main

//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
//"X","X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//
//
// 示例 2：
//
//
//输入：board = [["X"]]
//输出：[["X"]]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
//
//
// 👍 927 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte) {
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(board), len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= 0 && r < m && c >= 0 && c < n && board[r][c] == 'O' {
			board[r][c] = 'V'
			for _, dir := range dirs {
				nr, nc := r+dir[0], c+dir[1]
				dfs(nr, nc)
			}
		}
	}
	for r := 0; r < m; r++ {
		dfs(r, 0)
		dfs(r, n-1)
	}
	for c := 0; c < n; c++ {
		dfs(0, c)
		dfs(m-1, c)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'V' {
				board[r][c] = 'O'
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

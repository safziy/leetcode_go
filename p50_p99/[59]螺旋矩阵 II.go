package main

//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
//
// 👍 1003 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir[0], col+dir[1]; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4
			// 顺时针旋转至下一个方向
			dir = dirs[dirIdx]
		}
		row += dir[0]
		col += dir[1]
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)

package main

//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
// 
// 
//
// 
//
// 示例 1： 
// 
// 
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
// 
//
// 示例 2： 
// 
// 
//输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
//输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[0].length 
// 1 <= m, n <= 200 
// -2³¹ <= matrix[i][j] <= 2³¹ - 1 
// 
//
// 
//
// 进阶： 
//
// 
// 一个直观的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。 
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。 
// 你能想出一个仅使用常量空间的解决方案吗？ 
// 
//
// 👍 854 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowFlag := 1
	for c := 0; c < n; c++ {
		if matrix[0][c] == 0 {
			firstRowFlag = 0
		}
		for r := 1; r < m; r++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}

		}
	}
	for c := n-1;c>=0;c-- {
		for r := m-1; r > 0; r-- {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
		if firstRowFlag == 0 {
			matrix[0][c] = 0
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

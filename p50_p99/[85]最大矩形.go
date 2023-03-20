package main

//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
//
//
// 示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
//
//
// 示例 2：
//
//
//输入：matrix = []
//输出：0
//
//
// 示例 3：
//
//
//输入：matrix = [["0"]]
//输出：0
//
//
// 示例 4：
//
//
//输入：matrix = [["1"]]
//输出：1
//
//
// 示例 5：
//
//
//输入：matrix = [["0","0"]]
//输出：0
//
//
//
//
// 提示：
//
//
// rows == matrix.length
// cols == matrix[0].length
// 1 <= row, cols <= 200
// matrix[i][j] 为 '0' 或 '1'
//
//
// 👍 1479 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m, n := len(matrix), len(matrix[0])
	heights := make([][]int, m)
	for r, row := range matrix {
		heights[r] = make([]int, n)
		for c, cell := range row {
			if cell == '0' {
				continue
			}
			if r == 0 {
				heights[r][c] = 1
			} else {
				heights[r][c] = heights[r-1][c] + 1
			}
		}
	}
	result := 0
	for r := 0; r < m; r++ {
		left, right := make([]int, n), make([]int, n)
		stack := []int{}
		for c := 0; c < n; c++ {
			for len(stack) > 0 && heights[r][stack[len(stack)-1]] >= heights[r][c] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				left[c] = -1
			} else {
				left[c] = stack[len(stack)-1]
			}
			stack = append(stack, c)
		}
		stack = []int{}
		for c := n - 1; c >= 0; c-- {
			for len(stack) > 0 && heights[r][stack[len(stack)-1]] >= heights[r][c] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				right[c] = n
			} else {
				right[c] = stack[len(stack)-1]
			}
			stack = append(stack, c)
		}
		for c := 0; c < n; c++ {
			result = max(result, (right[c]-left[c]-1)*heights[r][c])
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

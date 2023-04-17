package main

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// 👍 1344 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	total := m * n
	result, idx := make([]int, total), 0
	left, right, top, bottom := 0, n-1, 0, m-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result[idx] = matrix[top][i]
			idx++
		}
		for i := top+1; i <= bottom; i++ {
			result[idx] = matrix[i][right]
			idx++
		}
		if left < right && top < bottom {
			for i := right-1; i >= left; i-- {
				result[idx] = matrix[bottom][i]
				idx++
			}
			for i := bottom-1; i > top; i-- {
				result[idx] = matrix[i][left]
				idx++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

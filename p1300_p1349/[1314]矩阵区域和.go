package p1300_p1349

//给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素
//mat[r][c] 的和：
//
//
// i - k <= r <= i + k,
// j - k <= c <= j + k 且
// (r, c) 在矩阵内。
//
//
//
//
// 示例 1：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
//输出：[[12,21,16],[27,45,33],[24,39,28]]
//
//
// 示例 2：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
//输出：[[45,45,45],[45,45,45],[45,45,45]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n, k <= 100
// 1 <= mat[i][j] <= 100
//
//
// 👍 162 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	preSum := make([][]int, m+1)
	for i := range preSum {
		preSum[i] = make([]int, n+1)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			preSum[r+1][c+1] = mat[r][c] + preSum[r][c+1] + preSum[r+1][c] - preSum[r][c]
		}
	}
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			up, down, left, right := r-k, r+k+1, c-k, c+k+1
			if up < 0 {
				up = 0
			}
			if left < 0 {
				left = 0
			}
			if down > m {
				down = m
			}
			if right > n {
				right = n
			}
			result[r][c] = preSum[down][right] - preSum[up][right] - preSum[down][left] + preSum[up][left]
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

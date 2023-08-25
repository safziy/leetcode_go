package main

//给你一个大小为 m x n 的矩阵 mat 和一个整数阈值 threshold。
//
// 请你返回元素总和小于或等于阈值的正方形区域的最大边长；如果没有这样的正方形区域，则返回 0 。
//
// 示例 1：
//
//
//
//
//输入：mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
//输出：2
//解释：总和小于或等于 4 的正方形的最大边长为 2，如图所示。
//
//
// 示例 2：
//
//
//输入：mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]],
//threshold = 1
//输出：0
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 300
// 0 <= mat[i][j] <= 10⁴
// 0 <= threshold <= 10⁵
//
//
// 👍 109 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}
	// 计算前缀和
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	sumHelper := func(x1, y1, x2, y2 int) int {
		return preSum[x2][y2] - preSum[x2][y1] - preSum[x1][y2] + preSum[x1][y1]
	}
	// 方法一: 二分查找
	check := func(mid int) bool {
		for i := 1; i+mid-1 <= m; i++ {
			for j := 1; j+mid-1 <= n; j++ {
				if sumHelper(i-1, j-1, i+mid-1, j+mid-1) <= threshold {
					return true
				}
			}
		}
		return false
	}
	left, right, res := 0, min(m, n), 0
	for left <= right {
		mid := left + (right-left)>>1
		if check(mid) {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

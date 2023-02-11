package main

//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
//
//
// 示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
//
//
// 示例 2：
//
//
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
//
//
// 示例 3：
//
//
//输入：matrix = [["0"]]
//输出：0
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'
//
//
// Related Topics 数组 动态规划 矩阵 👍 1254 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	max := 0
	for i := 1; i <= m; i++ {
		temp := 0
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] != '0' {
				v := min(min(dp[j], dp[j-1]), temp) + 1
				temp = dp[j]
				dp[j] = v
				if v > max {
					max = v
				}
			} else {
				dp[j] = 0
			}
		}
		//fmt.Println(dp)
	}
	return max * max
}

//leetcode submit region end(Prohibit modification and deletion)

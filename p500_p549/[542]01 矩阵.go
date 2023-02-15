package main

//给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
//
// 两个相邻元素间的距离为 1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
//输出：[[0,0,0],[0,1,0],[0,0,0]]
//
//
// 示例 2：
//
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
//输出：[[0,0,0],[0,1,0],[1,2,1]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10⁴
// 1 <= m * n <= 10⁴
// mat[i][j] is either 0 or 1.
// mat 中至少有一个 0
//
//
// Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 814 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(mat [][]int) [][]int {
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(mat), len(mat[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	queue := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		for _, dir := range dirs {
			r, c := queue[0][0]+dir[0], queue[0][1]+dir[1]
			if r >= 0 && r < m && c >= 0 && c < n && mat[r][c] == 1 && result[r][c] == 0 {
				result[r][c] = result[queue[0][0]][queue[0][1]] + 1
				queue = append(queue, [2]int{r, c})
			}
		}
		queue = queue[1:]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

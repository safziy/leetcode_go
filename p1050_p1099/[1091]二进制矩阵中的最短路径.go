package p1050_p1099

//给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
//
// 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求
//：
//
//
// 路径途经的所有单元格都的值都是 0 。
// 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
//
//
// 畅通路径的长度 是该路径途经的单元格总数。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,1],[1,0]]
//输出：2
//
//
// 示例 2：
//
//
//输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
//输出：4
//
//
// 示例 3：
//
//
//输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
//输出：-1
//
//
//
//
// 提示：
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 100
// grid[i][j] 为 0 或 1
//
//
// 👍 251 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	dirs := [][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	queue := [][2]int{}
	queue = append(queue, [2]int{0, 0})
	grid[0][0] = -1 // -x 表示遍历到该为止需要的步数为-x
	for len(queue) > 0 {
		if queue[0][0] == n-1 && queue[0][1] == n-1 {
			return -grid[queue[0][0]][queue[0][1]]
		}
		for _, dir := range dirs {
			r, c := queue[0][0]+dir[0], queue[0][1]+dir[1]
			if r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
				grid[r][c] = grid[queue[0][0]][queue[0][1]] - 1
			}
		}
		queue = queue[1:]
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

package p1200_p1249

//你还记得那条风靡全球的贪吃蛇吗？
//
// 我们在一个 n*n 的网格上构建了新的迷宫地图，蛇的长度为 2，也就是说它会占去两个单元格。蛇会从左上角（(0, 0) 和 (0, 1)）开始移动。我们用
// 0 表示空单元格，用 1 表示障碍物。蛇需要移动到迷宫的右下角（(n-1, n-2) 和 (n-1, n-1)）。
//
// 每次移动，蛇可以这样走：
//
//
// 如果没有障碍，则向右移动一个单元格。并仍然保持身体的水平／竖直状态。
// 如果没有障碍，则向下移动一个单元格。并仍然保持身体的水平／竖直状态。
// 如果它处于水平状态并且其下面的两个单元都是空的，就顺时针旋转 90 度。蛇从（(r, c)、(r, c+1)）移动到 （(r, c)、(r+1, c)）。
//
// 如果它处于竖直状态并且其右面的两个单元都是空的，就逆时针旋转 90 度。蛇从（(r, c)、(r+1, c)）移动到（(r, c)、(r, c+1)）。
//
//
//
// 返回蛇抵达目的地所需的最少移动次数。
//
// 如果无法到达目的地，请返回 -1。
//
//
//
// 示例 1：
//
//
//
// 输入：grid = [[0,0,0,0,0,1],
//               [1,1,0,0,1,0],
//               [0,0,0,0,1,1],
//               [0,0,1,0,1,0],
//               [0,1,1,0,0,0],
//               [0,1,1,0,0,0]]
//输出：11
//解释：
//一种可能的解决方案是 [右, 右, 顺时针旋转, 右, 下, 下, 下, 下, 逆时针旋转, 右, 下]。
//
//
// 示例 2：
//
// 输入：grid = [[0,0,1,1,1,1],
//               [0,0,0,0,1,1],
//               [1,1,0,0,0,1],
//               [1,1,1,0,0,1],
//               [1,1,1,0,0,1],
//               [1,1,1,0,0,0]]
//输出：9
//
//
//
//
// 提示：
//
//
// 2 <= n <= 100
// 0 <= grid[i][j] <= 1
// 蛇保证从空单元格开始出发。
//
//
// Related Topics 广度优先搜索 数组 矩阵 👍 122 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(grid [][]int) int {
	type pair struct {
		r, c int  // 行，列
		d    bool // 方向 true-右 false-下
		step int
	}
	key := func(r, c int, d bool) int {
		if d {
			return 1<<16 + r<<8 + c
		}
		return r<<8 + c
	}
	n := len(grid)
	visit := make(map[int]bool, 2*n)
	queue := make([]pair, 0, 2*n)
	queue = append(queue, pair{0, 1, true, 0})
	visit[key(0, 1, true)] = true
	for len(queue) > 0 {
		cur := queue[0]
		if cur.r == n-1 && cur.c == n-1 && cur.d {
			return cur.step
		}
		if cur.d {
			if cur.c+1 < n && grid[cur.r][cur.c+1] == 0 {
				k := key(cur.r, cur.c+1, true) // 右移
				if visit[k] == false {
					queue = append(queue, pair{cur.r, cur.c + 1, true, cur.step + 1})
					visit[k] = true
				}
			}
			if cur.r+1 < n && grid[cur.r+1][cur.c-1] == 0 && grid[cur.r+1][cur.c] == 0 {
				k := key(cur.r+1, cur.c, true) // 下移
				if visit[k] == false {
					queue = append(queue, pair{cur.r + 1, cur.c, true, cur.step + 1})
					visit[k] = true
				}
				k = key(cur.r+1, cur.c-1, false) //  顺时针旋转
				if visit[k] == false {
					queue = append(queue, pair{cur.r + 1, cur.c - 1, false, cur.step + 1})
					visit[k] = true
				}
			}
		} else {
			if cur.r+1 < n && grid[cur.r+1][cur.c] == 0 {
				k := key(cur.r+1, cur.c, false) // 下移
				if visit[k] == false {
					queue = append(queue, pair{cur.r + 1, cur.c, false, cur.step + 1})
					visit[k] = true
				}
			}
			if cur.c+1 < n && grid[cur.r-1][cur.c+1] == 0 && grid[cur.r][cur.c+1] == 0 {
				k := key(cur.r, cur.c+1, false) // 右移
				if visit[k] == false {
					queue = append(queue, pair{cur.r, cur.c + 1, false, cur.step + 1})
					visit[k] = true
				}
				k = key(cur.r-1, cur.c+1, true) //  逆时针旋转
				if visit[k] == false {
					queue = append(queue, pair{cur.r - 1, cur.c + 1, true, cur.step + 1})
					visit[k] = true
				}
			}
		}
		queue = queue[1:]
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

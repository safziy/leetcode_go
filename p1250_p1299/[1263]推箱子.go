package main

//「推箱子」是一款风靡全球的益智小游戏，玩家需要将箱子推到仓库中的目标位置。
//
// 游戏地图用大小为 m x n 的网格 grid 表示，其中每个元素可以是墙、地板或者是箱子。
//
// 现在你将作为玩家参与游戏，按规则将箱子 'B' 移动到目标位置 'T' ：
//
//
// 玩家用字符 'S' 表示，只要他在地板上，就可以在网格中向上、下、左、右四个方向移动。
// 地板用字符 '.' 表示，意味着可以自由行走。
// 墙用字符 '#' 表示，意味着障碍物，不能通行。
// 箱子仅有一个，用字符 'B' 表示。相应地，网格上有一个目标位置 'T'。
// 玩家需要站在箱子旁边，然后沿着箱子的方向进行移动，此时箱子会被移动到相邻的地板单元格。记作一次「推动」。
// 玩家无法越过箱子。
//
//
// 返回将箱子推到目标位置的最小 推动 次数，如果无法做到，请返回 -1。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [["#","#","#","#","#","#"],
//             ["#","T","#","#","#","#"],
//             ["#",".",".","B",".","#"],
//             ["#",".","#","#",".","#"],
//             ["#",".",".",".","S","#"],
//             ["#","#","#","#","#","#"]]
//输出：3
//解释：我们只需要返回推箱子的次数。
//
// 示例 2：
//
//
//输入：grid = [["#","#","#","#","#","#"],
//             ["#","T","#","#","#","#"],
//             ["#",".",".","B",".","#"],
//             ["#","#","#","#",".","#"],
//             ["#",".",".",".","S","#"],
//             ["#","#","#","#","#","#"]]
//输出：-1
//
//
// 示例 3：
//
//
//输入：grid = [["#","#","#","#","#","#"],
//             ["#","T",".",".","#","#"],
//             ["#",".","#","B",".","#"],
//             ["#",".",".",".",".","#"],
//             ["#",".",".",".","S","#"],
//             ["#","#","#","#","#","#"]]
//输出：5
//解释：向下、向左、向左、向上再向上。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 20
// grid 仅包含字符 '.', '#', 'S' , 'T', 以及 'B'。
// grid 中 'S', 'B' 和 'T' 各只能出现一个。
//
//
// 👍 122 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	hash := func(r, c, d int) int {
		return r<<10 | c<<5 | d
	}
	check := func(r, c int) bool {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == '#' {
			return false
		}
		return true
	}
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// 当前箱子在box处判断玩家是否可以从start是否可以到达target
	reachable := func(start, target, box [2]int) bool {
		visit := map[int]bool{}
		queue := [][2]int{}
		queue = append(queue, start)
		visit[hash(start[0], start[1], 0)] = true
		for len(queue) > 0 {
			cur := queue[0]
			if cur[0] == target[0] && cur[1] == target[1] {
				return true
			}
			for _, dir := range dirs {
				nr, nc := cur[0]+dir[0], cur[1]+dir[1]
				h := hash(nr, nc, 0)
				if check(nr, nc) && visit[h] == false && (nr != box[0] || nc != box[1]) {
					queue = append(queue, [2]int{nr, nc})
					visit[h] = true
				}
			}
			queue = queue[1:]
		}
		return false
	}
	// 找到起点，终点，箱子的位置
	var startPos, targetPos, boxPos [2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 'S':
				startPos = [2]int{i, j}
			case 'T':
				targetPos = [2]int{i, j}
			case 'B':
				boxPos = [2]int{i, j}
			}
		}
	}
	// 广度优先搜索,队列中存放的是 {箱子的位置，玩家的方向，推箱子的次数}
	queue := [][4]int{}
	visit := map[int]bool{}
	// 遍历可以推箱子的方向
	for i, dir := range dirs {
		pr, pc := boxPos[0]+dir[0], boxPos[1]+dir[1]
		if check(pr, pc) && reachable(startPos, [2]int{pr, pc}, boxPos) {
			queue = append(queue, [4]int{boxPos[0], boxPos[1], i, 0})
			visit[hash(boxPos[0], boxPos[1], i)] = true
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		if cur[0] == targetPos[0] && cur[1] == targetPos[1] {
			return cur[3]
		}
		boxPos = [2]int{cur[0], cur[1]}
		playerPos := [2]int{cur[0] + dirs[cur[2]][0], cur[1] + dirs[cur[2]][1]}
		// 遍历可以推箱子的方向
		for i, dir := range dirs {
			// 玩家需要站的位置和箱子需要推到的位置
			pr, pc, tr, tc := cur[0]+dir[0], cur[1]+dir[1], cur[0]-dir[0], cur[1]-dir[1]
			h := hash(tr, tc, i)
			if check(pr, pc) && check(tr, tc) && visit[h] == false && reachable(playerPos, [2]int{pr, pc}, boxPos) {
				queue = append(queue, [4]int{tr, tc, i, cur[3] + 1})
				visit[h] = true
			}
		}
		queue = queue[1:]
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

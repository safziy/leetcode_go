package main

//在一个有向图中，节点分别标记为 0, 1, ..., n-1。图中每条边为红色或者蓝色，且存在自环或平行边。
//
// red_edges 中的每一个 [i, j] 对表示从节点 i 到节点 j 的红色有向边。类似地，blue_edges 中的每一个 [i, j] 对表示从
//节点 i 到节点 j 的蓝色有向边。
//
// 返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，
//那么 answer[x] = -1。
//
//
//
// 示例 1：
//
//
//输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
//输出：[0,1,-1]
//
//
// 示例 2：
//
//
//输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
//输出：[0,1,-1]
//
//
// 示例 3：
//
//
//输入：n = 3, red_edges = [[1,0]], blue_edges = [[2,1]]
//输出：[0,-1,-1]
//
//
// 示例 4：
//
//
//输入：n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
//输出：[0,1,2]
//
//
// 示例 5：
//
//
//输入：n = 3, red_edges = [[0,1],[0,2]], blue_edges = [[1,0]]
//输出：[0,1,1]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// red_edges.length <= 400
// blue_edges.length <= 400
// red_edges[i].length == blue_edges[i].length == 2
// 0 <= red_edges[i][j], blue_edges[i][j] < n
//
//
// Related Topics 广度优先搜索 图 👍 187 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	g := make([][]pair, n)
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	vis := make([][2]bool, n)
	vis[0] = [2]bool{true, true}
	q := []pair{{0, 0}, {0, 1}}
	for level := 0; len(q) > 0; level++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x := p.x
			if dis[x] < 0 {
				dis[x] = level
			}
			for _, to := range g[x] {
				if to.color != p.color && !vis[to.x][to.color] {
					vis[to.x][to.color] = true
					q = append(q, to)
				}
			}
		}
	}
	return dis
}

//leetcode submit region end(Prohibit modification and deletion)

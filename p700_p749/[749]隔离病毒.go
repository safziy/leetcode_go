package main

//病毒扩散得很快，现在你的任务是尽可能地通过安装防火墙来隔离病毒。
//
// 假设世界由 m x n 的二维矩阵 isInfected 组成， isInfected[i][j] == 0 表示该区域未感染病毒，而
//isInfected[i][j] == 1 表示该区域已感染病毒。可以在任意 2 个相邻单元之间的共享边界上安装一个防火墙（并且只有一个防火墙）。
//
// 每天晚上，病毒会从被感染区域向相邻未感染区域扩散，除非被防火墙隔离。现由于资源有限，每天你只能安装一系列防火墙来隔离其中一个被病毒感染的区域（一个区域或连
//续的一片区域），且该感染区域对未感染区域的威胁最大且 保证唯一 。
//
// 你需要努力使得最后有部分区域不被病毒感染，如果可以成功，那么返回需要使用的防火墙个数; 如果无法实现，则返回在世界被病毒全部感染时已安装的防火墙个数。
//
//
//
// 示例 1：
//
//
//
//
//输入: isInfected = [[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0
//,0,0,0,0,0]]
//输出: 10
//解释:一共有两块被病毒感染的区域。
//在第一天，添加 5 墙隔离病毒区域的左侧。病毒传播后的状态是:
//
//第二天，在右侧添加 5 个墙来隔离病毒区域。此时病毒已经被完全控制住了。
//
//
//
// 示例 2：
//
//
//
//
//输入: isInfected = [[1,1,1],[1,0,1],[1,1,1]]
//输出: 4
//解释: 虽然只保存了一个小区域，但却有四面墙。
//注意，防火墙只建立在两个不同区域的共享边界上。
//
//
// 示例 3:
//
//
//输入: isInfected = [[1,1,1,0,0,0,0,0,0],[1,0,1,0,1,1,1,1,1],[1,1,1,0,0,0,0,0,0]]
//
//输出: 13
//解释: 在隔离右边感染区域后，隔离左边病毒区域只需要 2 个防火墙。
//
//
//
//
// 提示:
//
//
// m == isInfected.length
// n == isInfected[i].length
// 1 <= m, n <= 50
// isInfected[i][j] is either 0 or 1
// 在整个描述的过程中，总有一个相邻的病毒区域，它将在下一轮 严格地感染更多未受污染的方块
//
//
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 模拟 👍 134 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func containVirus(isInfected [][]int) int {
	type pair struct{ r, c int }
	m, n := len(isInfected), len(isInfected[0])
	result := 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for {
		// areaIdx:区域编号 firewalls[i]:第i个区域需要的防火墙数量 neighbors[i]:第i个区域下一天会感染的方块位置
		areaIdx, firewalls, neighbors := 0, []int{}, []map[pair]struct{}{}
		// areaIdx-区域编号
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				if isInfected[r][c] != 1 {
					continue
				}
				areaIdx++
				// firewall:需要的防火墙数量 neighbor:下一步感染的地方
				firewall, neighbor := 0, map[pair]struct{}{}
				queue := []pair{}
				queue = append(queue, pair{r, c})
				isInfected[r][c] = -areaIdx // 代表已经访问过
				for len(queue) > 0 {
					curR, curC := queue[0].r, queue[0].c
					queue = queue[1:]
					for _, dir := range dirs {
						next := pair{curR + dir[0], curC + dir[1]}
						if next.r >= 0 && next.r < m && next.c >= 0 && next.c < n {
							if isInfected[next.r][next.c] == 1 {
								isInfected[next.r][next.c] = -areaIdx
								queue = append(queue, next)
							} else if isInfected[next.r][next.c] == 0 {
								neighbor[next] = struct{}{}
								firewall++
							}
						}
					}
				}
				firewalls = append(firewalls, firewall)
				neighbors = append(neighbors, neighbor)
			}
		}
		// 没有可以感染的
		if len(neighbors) == 0 {
			break
		}
		// 找到威胁最大的
		maxIdx := 0
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[maxIdx]) {
				maxIdx = i
			}
		}
		// 加上需要安装的防火墙数量
		result += firewalls[maxIdx]
		// 如果只剩下一块区域，则下一次感染完就结束了
		if len(neighbors) == 1 {
			break
		}
		// 将之前已经访问过的标志删除
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				if isInfected[r][c] < 0 {
					if isInfected[r][c] == -maxIdx-1 {
						isInfected[r][c] = 2 // 2 标志为已经被防火墙包围
					} else {
						isInfected[r][c] = 1
					}
				}
			}
		}
		// 将所有的邻居都感染掉
		for i, neighbor := range neighbors {
			if i != maxIdx {
				for p := range neighbor {
					isInfected[p.r][p.c] = 1
				}
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

package main

//n 座城市和一些连接这些城市的道路 roads 共同组成一个基础设施网络。每个 roads[i] = [ai, bi] 都表示在城市 ai 和 bi 之间有
//一条双向道路。
//
// 两座不同城市构成的 城市对 的 网络秩 定义为：与这两座城市 直接 相连的道路总数。如果存在一条道路直接连接这两座城市，则这条道路只计算 一次 。
//
// 整个基础设施网络的 最大网络秩 是所有不同城市对中的 最大网络秩 。
//
// 给你整数 n 和数组 roads，返回整个基础设施网络的 最大网络秩 。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
//输出：4
//解释：城市 0 和 1 的网络秩是 4，因为共有 4 条道路与城市 0 或 1 相连。位于 0 和 1 之间的道路只计算一次。
//
//
// 示例 2：
//
//
//
//
//输入：n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
//输出：5
//解释：共有 5 条道路与城市 1 或 2 相连。
//
//
// 示例 3：
//
//
//输入：n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
//输出：5
//解释：2 和 5 的网络秩为 5，注意并非所有的城市都需要连接起来。
//
//
//
//
// 提示：
//
//
// 2 <= n <= 100
// 0 <= roads.length <= n * (n - 1) / 2
// roads[i].length == 2
// 0 <= ai, bi <= n-1
// ai != bi
// 每对城市之间 最多只有一条 道路相连
//
//
// 👍 69 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maximalNetworkRank(n int, roads [][]int) int {
	connect := make([][]bool, n)
	for i := range connect {
		connect[i] = make([]bool, n)
	}
	degree := make([]int, n)
	for _, road := range roads {
		connect[road[0]][road[1]] = true
		connect[road[1]][road[0]] = true
		degree[road[0]]++
		degree[road[1]]++
	}
	firstMax, secondMax, firstArr, secondArr := -1, -2, []int{}, []int{}
	for i := 0; i < n; i++ {
		if degree[i] > firstMax {
			firstMax, secondMax = degree[i], firstMax
			firstArr, secondArr = []int{}, firstArr
			firstArr = append(firstArr, i)
		} else if degree[i] == firstMax {
			firstArr = append(firstArr, i)
		} else if degree[i] > secondMax {
			secondMax, secondArr = degree[i], []int{}
			secondArr = append(secondArr, i)
		} else if degree[i] == secondMax {
			secondArr = append(secondArr, i)
		}
	}
	if len(firstArr) == 1 {
		f := firstArr[0]
		for _, s := range secondArr {
			if !connect[f][s] {
				return firstMax + secondMax
			}
		}
		return firstMax + secondMax - 1
	} else {
		if len(firstArr)*(len(firstArr)-1)/2 > len(roads) {
			return firstMax * 2
		}
		for _, f1 := range firstArr {
			for _, f2 := range firstArr {
				if f1 != f2 && !connect[f1][f2] {
					return firstMax * 2
				}
			}
		}
		return firstMax*2 - 1
	}
}

//leetcode submit region end(Prohibit modification and deletion)

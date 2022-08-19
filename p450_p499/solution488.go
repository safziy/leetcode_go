package p450_p499

/*
488. 祖玛游戏
你正在参与祖玛游戏的一个变种。

在这个祖玛游戏变体中，桌面上有 一排 彩球，每个球的颜色可能是：红色 'R'、黄色 'Y'、蓝色 'B'、绿色 'G' 或白色 'W' 。你的手中也有一些彩球。

你的目标是 清空 桌面上所有的球。每一回合：

从你手上的彩球中选出 任意一颗 ，然后将其插入桌面上那一排球中：两球之间或这一排球的任一端。
接着，如果有出现 三个或者三个以上 且 颜色相同 的球相连的话，就把它们移除掉。
如果这种移除操作同样导致出现三个或者三个以上且颜色相同的球相连，则可以继续移除这些球，直到不再满足移除条件。
如果桌面上所有球都被移除，则认为你赢得本场游戏。
重复这个过程，直到你赢了游戏或者手中没有更多的球。
给你一个字符串 board ，表示桌面上最开始的那排球。另给你一个字符串 hand ，表示手里的彩球。请你按上述操作步骤移除掉桌上所有球，计算并返回所需的 最少 球数。如果不能移除桌上所有的球，返回 -1 。

提示：

1 <= board.length <= 16
1 <= hand.length <= 5
board 和 hand 由字符 'R'、'Y'、'B'、'G' 和 'W' 组成
桌面上一开始的球中，不会有三个及三个以上颜色相同且连着的球

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zuma-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findMinStep(board string, hand string) int {
	convertMap := map[byte]byte{'R': 0, 'Y': 1, 'B': 2, 'G': 3, 'W': 4} // 先把手上的球都用数组表示
	convertString := "RYBGW"                                            // 数组转换成球
	searched := make(map[string]struct{})
	// dfs已遍历过的innerBoard
	// 小bug 没有保存当时的数组状态     可以把数组状态转换成string 追加在 innerBoard 后面
	result := int(^uint(0) >> 1) // int最大值
	handSlice := make([]int, 5)  // 各色球数量的数组

	// 初始化球数量
	for i := range hand {
		handSlice[convertMap[hand[i]]]++
	}

	// 深度优先搜索遍历
	var dfs func(string, int)
	dfs = func(innerBoard string, need int) {
		if len(innerBoard) == 0 && need < result {
			result = need // 祖玛列表为空时球数量的最小值
		}
		if _, ok := searched[innerBoard]; ok || need >= result {
			return // 剪枝， 遍历时如果需要的数量大于等于当前最优解 或者 已遍历过当前状态
		}

		for i := range innerBoard { // 回溯 把当前所有可插入为止插入一遍再进行判断
			for j := range handSlice { // 遍历球数组
				if handSlice[j] == 0 {
					continue
				}
				handSlice[j]--
				dfs(eliminate(innerBoard[:i]+convertString[j:j+1]+innerBoard[i:]), need+1)
				handSlice[j]++
			}
		}
		searched[innerBoard] = struct{}{}
	}
	dfs(board, 0)

	if result == int(^uint(0)>>1) {
		return -1
	} else {
		return result
	}
}

// 根据当前状态清球 应该可优化 但懒的了
func eliminate(board string) string {
	flag := true // 是否需要继续遍历
	for flag && len(board) > 0 {
		flag = false
		count := 1
		length := len(board)
		for i := 1; i < length; i++ {
			if board[i] == board[i-1] {
				count++
			} else {
				if count >= 3 {
					board = board[:i-count] + board[i:]
					flag = true
					break
				}
				count = 1
			}
		}
		if count >= 3 {
			board = board[:length-count]
			flag = true
		}
	}
	return board
}

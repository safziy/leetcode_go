package main

//一个 n x n 的二维网络 board 仅由 0 和 1 组成 。每次移动，你能任意交换两列或是两行的位置。
//
// 返回 将这个矩阵变为 “棋盘” 所需的最小移动次数 。如果不存在可行的变换，输出 -1。 
//
// “棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
//输出: 2
//解释:一种可行的变换方式如下，从左到右：
//第一次移动交换了第一列和第二列。
//第二次移动交换了第二行和第三行。
// 
//
// 示例 2: 
//
// 
//
// 
//输入: board = [[0, 1], [1, 0]]
//输出: 0
//解释: 注意左上角的格值为0时也是合法的棋盘，也是合法的棋盘.
// 
//
// 示例 3: 
//
// 
//
// 
//输入: board = [[1, 0], [1, 0]]
//输出: -1
//解释: 任意的变换都不能使这个输入变为合法的棋盘。
// 
//
// 
//
// 提示： 
//
// 
// n == board.length 
// n == board[i].length 
// 2 <= n <= 30 
// board[i][j] 将只包含 0或 1 
// 
//
// Related Topics 位运算 数组 数学 矩阵 👍 140 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func movesToChessboard(board [][]int) int {
	n := len(board)
	t, half := 1<<n-1, n/2              // 二进制全为1的数字，n的一半
	rowSum, rowCnt, rowRight := 0, 0, 0 // 第一行二进制的和，数字1的个数，1处在正确位置的数量(按1010来排列)
	colSum, colCnt, colRight := 0, 0, 0 // 第一列二进制的和，数字1的个数，1处在正确位置的数量(按1010来排列)
	for i := 0; i < n; i++ {
		if board[0][i] == 1 {
			rowSum |= 1 << i
			rowCnt++
			if i%2 == 0 {
				rowRight++
			}
		}
		if board[i][0] == 1 {
			colSum |= 1 << i
			colCnt++
			if i%2 == 0 {
				colRight++
			}
		}
	}
	// n为偶数，如果1的个数不是一半,则说明不能变换
	if n%2 == 0 && (rowCnt != half || colCnt != half) {
		return -1
	}
	// n不为偶数，如果1的个数不是一半或者一半加1,则说明不能变换
	if n%2 != 0 && ((rowCnt != half && rowCnt != half+1) || (colCnt != half && colCnt != half+1)) {
		return -1
	}
	for i := 1; i < n; i++ {
		rowTemp, colTemp := 0, 0
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				rowTemp |= 1 << j
			}
			if board[j][i] == 1 {
				colTemp |= 1 << j
			}
		}
		// 如果其他行不和第一行一样或者互补，则不能变换
		if rowTemp != rowSum && rowTemp^rowSum != t {
			return -1
		}
		// 如果其他列不和第一列一样或者互补，则不能变换
		if colTemp != colSum && colTemp^colSum != t {
			return -1
		}
	}
	rowStep, colStep := 0, 0
	if n%2 == 0 {
		rowStep = min(rowRight, half-rowRight)
		colStep = min(colRight, half-colRight)
	} else {
		if rowCnt == half {
			rowStep = rowRight
		} else {
			rowStep = half - rowRight + 1
		}
		if colCnt == half {
			colStep = colRight
		} else {
			colStep = half - colRight + 1
		}
	}
	return rowStep + colStep
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)

package main

//我们从一块字母板上的位置 (0, 0) 出发，该坐标对应的字符为 board[0][0]。
//
// 在本题里，字母板为board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"]，如下所示。
//
//
//
// 我们可以按下面的指令规则行动：
//
//
// 如果方格存在，'U' 意味着将我们的位置上移一行；
// 如果方格存在，'D' 意味着将我们的位置下移一行；
// 如果方格存在，'L' 意味着将我们的位置左移一列；
// 如果方格存在，'R' 意味着将我们的位置右移一列；
// '!' 会把在我们当前位置 (r, c) 的字符 board[r][c] 添加到答案中。
//
//
// （注意，字母板上只存在有字母的位置。）
//
// 返回指令序列，用最小的行动次数让答案和目标 target 相同。你可以返回任何达成目标的路径。
//
//
//
// 示例 1：
//
//
//输入：target = "leet"
//输出："DDR!UURRR!!DDD!"
//
//
// 示例 2：
//
//
//输入：target = "code"
//输出："RR!DDRR!UUL!R!"
//
//
//
//
// 提示：
//
//
// 1 <= target.length <= 100
// target 仅含有小写英文字母。
//
//
// Related Topics 哈希表 字符串 👍 70 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func alphabetBoardPath(target string) string {
	move := func(diff int32, isHorizontal bool) []byte {
		var ch byte = ' '
		if diff < 0 && isHorizontal {
			ch = 'L'
			diff = -diff
		} else if diff < 0 && !isHorizontal {
			ch = 'U'
			diff = -diff
		} else if diff >= 0 && isHorizontal {
			ch = 'R'
		} else {
			ch = 'D'
		}
		step := make([]byte, diff)
		for i := 0; i < int(diff); i++ {
			step[i] = ch
		}
		return step
	}
	result := make([]byte, 0, 500)
	var curRow, curCol int32 = 0, 0
	for _, ch := range target {
		d := ch - 'a'
		row, col := d/5, d%5
		rowDiff, colDiff := row-curRow, col-curCol
		if ch == 'z' { // 如果目标是'z' 则先左右移动
			result = append(result, move(colDiff, true)...)
			result = append(result, move(rowDiff, false)...)
		} else { // 其他情况都先上下移动
			result = append(result, move(rowDiff, false)...)
			result = append(result, move(colDiff, true)...)
		}
		result = append(result, '!')
		curRow, curCol = row, col
	}
	return string(result)
}

//leetcode submit region end(Prohibit modification and deletion)

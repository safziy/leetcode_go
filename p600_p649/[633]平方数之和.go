package main

import "math"

//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a² + b² = c 。
//
// 
//
// 示例 1： 
//
// 
//输入：c = 5
//输出：true
//解释：1 * 1 + 2 * 2 = 5
// 
//
// 示例 2： 
//
// 
//输入：c = 3
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 0 <= c <= 2³¹ - 1 
// 
//
// 👍 424 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func judgeSquareSum(c int) bool {
	a, b := 0, int(math.Sqrt(float64(c)))
	for a <= b {
		sum := a * a + b* b
		if sum == c {
			return true
		} else if sum > c {
			b --
		} else {
			a ++
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)

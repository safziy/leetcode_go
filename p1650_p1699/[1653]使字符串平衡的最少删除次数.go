package main

//给你一个字符串 s ，它仅包含字符 'a' 和 'b' 。
//
// 你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a'
//，此时认为 s 是 平衡 的。
//
// 请你返回使 s 平衡 的 最少 删除次数。
//
//
//
// 示例 1：
//
//
//输入：s = "aababbab"
//输出：2
//解释：你可以选择以下任意一种方案：
//下标从 0 开始，删除第 2 和第 6 个字符（"aababbab" -> "aaabbb"），
//下标从 0 开始，删除第 3 和第 6 个字符（"aababbab" -> "aabbbb"）。
//
//
// 示例 2：
//
//
//输入：s = "bbaaaaabb"
//输出：2
//解释：唯一的最优解是删除最前面两个字符。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s[i] 要么是 'a' 要么是 'b' 。
//
//
// 👍 87 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minimumDeletions(s string) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// aCnt 为从右边数有多少个a bCnt为从左边数有多少个b
	aCnt, bCnt := 0, 0
	for _, ch := range s {
		if ch == 'a' {
			aCnt++
		}
	}
	result := aCnt
	for _, ch := range s {
		if ch == 'a' {
			aCnt--
		} else {
			bCnt++
		}
		result = min(result, aCnt+bCnt)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

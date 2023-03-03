package main

//给定两个单词 word1 和
// word2 ，返回使得
// word1 和
// word2 相同所需的最小步数。
//
// 每步 可以删除任意一个字符串中的一个字符。
//
//
//
// 示例 1：
//
//
//输入: word1 = "sea", word2 = "eat"
//输出: 2
//解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
//
//
// 示例 2:
//
//
//输入：word1 = "leetcode", word2 = "etco"
//输出：4
//
//
//
//
// 提示：
//
//
//
// 1 <= word1.length, word2.length <= 500
// word1 和 word2 只包含小写英文字母
//
//
// 👍 539 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)

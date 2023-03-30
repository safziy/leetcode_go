package p1050_p1099

//给出两个字符串 str1 和 str2，返回同时以 str1 和 str2 作为子序列的最短字符串。如果答案不止一个，则可以返回满足条件的任意一个答案。
//
// （如果从字符串 T 中删除一些字符（也可能不删除，并且选出的这些字符可以位于 T 中的 任意位置），可以得到字符串 S，那么 S 就是 T 的子序列）
//
//
//
// 示例：
//
// 输入：str1 = "abac", str2 = "cab"
//输出："cabac"
//解释：
//str1 = "abac" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 的第一个 "c"得到 "abac"。
//str2 = "cab" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 末尾的 "ac" 得到 "cab"。
//最终我们给出的答案是满足上述属性的最短字符串。
//
//
//
//
// 提示：
//
//
// 1 <= str1.length, str2.length <= 1000
// str1 和 str2 都由小写英文字母组成。
//
//
// 👍 191 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func shortestCommonSupersequence(str1 string, str2 string) string {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n1, n2 := len(str1), len(str2)
	// 先根据动态规划计算最小长度
	// dp[i][j] 表示str1[0:i]和str1[0:j]公共超序列的长度
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n2; j++ {
		dp[0][j] = j
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	// 根据动态规划过程构造字符串
	result := make([]byte, dp[n1][n2])
	i, j, idx := n1, n2, dp[n1][n2]
	for i > 0 && j > 0 {
		idx--
		if str1[i-1] == str2[j-1] {
			i--
			j--
			result[idx] = str1[i]
		} else if dp[i][j] == dp[i-1][j]+1 {
			i--
			result[idx] = str1[i]
		} else if dp[i][j] == dp[i][j-1]+1 {
			j--
			result[idx] = str2[j]
		}
	}
	for i > 0 {
		idx--
		i--
		result[idx] = str1[i]
	}
	for j > 0 {
		idx--
		j--
		result[idx] = str2[j]
	}
	return string(result)
}

//leetcode submit region end(Prohibit modification and deletion)

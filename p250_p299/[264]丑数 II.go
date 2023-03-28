package main

//给你一个整数 n ，请你找出并返回第 n 个 丑数 。
//
// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 10
//输出：12
//解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：1
//解释：1 通常被视为丑数。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 1690 
// 
//
// 👍 1048 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num2, num3, num5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(num2, min(num3, num5))
		if dp[i] == num2 {
			p2++
		}
		if dp[i] == num3 {
			p3++
		}
		if dp[i] == num5 {
			p5++
		}
	}
	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

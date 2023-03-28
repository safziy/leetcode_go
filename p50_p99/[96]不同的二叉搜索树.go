package main
//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。 
//
// 
//
// 示例 1： 
// 
// 
//输入：n = 3
//输出：5
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 19 
// 
//
// 👍 2155 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	// 方法一： 动态规划
	// dp[i]长度为 i 的序列能构成的不同二叉搜索树的个数
	//dp := make([]int, n +1)
	//dp[0], dp[1] = 1, 1
	//for i := 2; i <= n; i++ {
	//	for j := 1; j <= i; j++ {
	//		dp[i] += dp[j-1] * dp[i-j]
	//	}
	//}
	//return dp[n]
	// 方法二： 数学-卡塔兰数？
	result := 1
	for i := 0; i < n; i++ {
		result = result * 2 * (2*i + 1) / (i + 2)
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

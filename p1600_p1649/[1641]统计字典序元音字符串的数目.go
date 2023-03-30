package main

//给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。
//
// 字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 1
//输出：5
//解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]
// 
//
// 示例 2： 
//
// 
//输入：n = 2
//输出：15
//解释：仅由元音组成的 15 个字典序字符串为
//["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
//注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后
// 
//
// 示例 3： 
//
// 
//输入：n = 33
//输出：66045
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 50 
// 
//
// 👍 106 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func countVowelStrings(n int) int {
	// 方法一：动态规划
	// dp[i]代表第i个字符开头的字符串数量
	// 添加一个新字符时 a可以添加到所有字符开头 e只能添加到 e，i，o，u前，依次类推
	// 即 dp[4]=dp[4]+dp[3]+dp[2]+dp[1]+dp[0] dp[3]=dp[3]+dp[2]+dp[1]+dp[0] dp[2]=dp[2]+dp[1]+dp[0] ...
	//dp := make([]int, 5)
	//dp[0] = 1
	//for i := 0; i <= n; i++ {
	//	for j := 1; j < 5; j++ {
	//		dp[j] = dp[j] + dp[j-1]
	//	}
	//}
	//return dp[4]

	// 方法二： 数学
	// 将n个字符分成5类，单类数量可以为空，即在n个字符中插入4个隔板
	// 先考虑单类数量不能为空的情况，即在加5个字符，每类默认放一个
	// 则问题变成总字符变成 n + 5，插入4个隔板，不能为空，则答案为 C(n+5-1, 4)
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}

//leetcode submit region end(Prohibit modification and deletion)

package main
//给定一个包含大写字母和小写字母的字符串
// s ，返回 通过这些字母构造成的 最长的回文串 。 
//
// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。 
//
// 
//
// 示例 1: 
//
// 
//输入:s = "abccccdd"
//输出:7
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
// 
//
// 示例 2: 
//
// 
//输入:s = "a"
//输出:1
// 
//
// 示例 3： 
//
// 
//输入:s = "aaaaaccc"
//输出:7 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length <= 2000 
// s 只由小写 和/或 大写英文字母组成 
// 
//
// 👍 517 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	cntMap := make(map[int32]int, 52)
	for _, ch := range s {
		cntMap[ch] ++
	}
	result := 0
	for _, cnt := range cntMap {
		result += cnt / 2 * 2
		if cnt & 1 == 1 && result & 1 == 0 {
			result ++
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

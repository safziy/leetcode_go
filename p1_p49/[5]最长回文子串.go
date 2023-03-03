package main

//给你一个字符串 s，找到 s 中最长的回文子串。
//
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// 👍 6191 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	n := len(s)
	expand := func(left, right int) (int, int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left, right = left-1, right+1
		}
		return left + 1, right - 1
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		left1, right1 := expand(i, i)
		left2, right2 := expand(i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

//leetcode submit region end(Prohibit modification and deletion)

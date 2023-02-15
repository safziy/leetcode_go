package main

import (
	"unicode"
)

//给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
//
// 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
//
//
//
// 示例 1：
//
//
//输入：s = "a1b2"
//输出：["a1b2", "a1B2", "A1b2", "A1B2"]
//
//
// 示例 2:
//
//
//输入: s = "3z4"
//输出: ["3z4","3Z4"]
//
//
//
//
// 提示:
//
//
// 1 <= s.length <= 12
// s 由小写英文字母、大写英文字母和数字组成
//
//
// Related Topics 位运算 字符串 回溯 👍 514 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func letterCasePermutation(s string) []string {
	n, cnt := len(s), 0
	for i := 0; i < n; i++ {
		if unicode.IsLetter(rune(s[i])) {
			cnt++
		}
	}
	i := 0
	result := make([]string, 1<<cnt)
	comb := make([]byte, n)
	var recursion func(idx int)
	recursion = func(idx int) {
		if idx == n {
			result[i] = string(comb)
			i++
			return
		}
		comb[idx] = s[idx]
		recursion(idx + 1)
		ch := rune(s[idx])
		if unicode.IsLower(ch) {
			comb[idx] = byte(unicode.ToUpper(ch))
			recursion(idx + 1)
		} else if unicode.IsUpper(ch) {
			comb[idx] = byte(unicode.ToLower(ch))
			recursion(idx + 1)
		}
	}
	recursion(0)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

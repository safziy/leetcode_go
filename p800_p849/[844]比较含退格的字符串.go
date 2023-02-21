package p800_p849

//给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
//
// 注意：如果对空文本输入退格字符，文本继续为空。
//
//
//
// 示例 1：
//
//
//输入：s = "ab#c", t = "ad#c"
//输出：true
//解释：s 和 t 都会变成 "ac"。
//
//
// 示例 2：
//
//
//输入：s = "ab##", t = "c#d#"
//输出：true
//解释：s 和 t 都会变成 ""。
//
//
// 示例 3：
//
//
//输入：s = "a#c", t = "b"
//输出：false
//解释：s 会变成 "c"，但 t 仍然是 "b"。
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 200
// s 和 t 只含有小写字母以及字符 '#'
//
//
//
//
// 进阶：
//
//
// 你可以用 O(n) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？
//
//
// 👍 584 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	sIdx, tIdx, sBack, tBack := sLen-1, tLen-1, 0, 0
	for sIdx >= 0 || tIdx >= 0 {
		for sIdx >= 0 && (s[sIdx] == '#' || sBack > 0) {
			if s[sIdx] == '#' {
				sBack++
			} else {
				sBack--
			}
			sIdx--
		}
		for tIdx >= 0 && (t[tIdx] == '#' || tBack > 0) {
			if t[tIdx] == '#' {
				tBack++
			} else {
				tBack--
			}
			tIdx--
		}
		if tIdx >= 0 && sIdx >= 0 {
			if s[sIdx] != t[tIdx] {
				return false
			}
			tIdx--
			sIdx--
		} else if tIdx < 0 && sIdx < 0 {
			return true
		} else if tIdx < 0 && s[sIdx] != '#' {
			return false
		} else if sIdx < 0 && t[tIdx] != '#' {
			return false
		}
	}
	return tIdx < 0 && sIdx < 0
}

//leetcode submit region end(Prohibit modification and deletion)

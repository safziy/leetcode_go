package main

//有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
//
// 假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
//
//
//
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
//
// 你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
//
// 请返回待替换子串的最小可能长度。
//
// 如果原字符串自身就是一个平衡字符串，则返回 0。
//
//
//
// 示例 1：
//
//
//输入：s = "QWER"
//输出：0
//解释：s 已经是平衡的了。
//
// 示例 2：
//
//
//输入：s = "QQWE"
//输出：1
//解释：我们需要把一个 'Q' 替换成 'R'，这样得到的 "RQWE" (或 "QRWE") 是平衡的。
//
//
// 示例 3：
//
//
//输入：s = "QQQW"
//输出：2
//解释：我们可以把前面的 "QQ" 替换成 "ER"。
//
//
// 示例 4：
//
//
//输入：s = "QQQQ"
//输出：3
//解释：我们可以替换后 3 个 'Q'，使 s = "QWER"。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10^5
// s.length 是 4 的倍数
// s 中只含有 'Q', 'W', 'E', 'R' 四种字符
//
//
// Related Topics 字符串 滑动窗口 👍 140 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func balancedString(s string) int {
	n := len(s)
	effectiveCnt := n / 4
	cnt := [4]int{}
	check := func() bool {
		if cnt[0] > effectiveCnt || cnt[1] > effectiveCnt || cnt[2] > effectiveCnt || cnt[3] > effectiveCnt {
			return false
		}
		return true
	}
	charIdx := func(ch uint8) int {
		if ch == 'Q' {
			return 0
		} else if ch == 'W' {
			return 1
		} else if ch == 'E' {
			return 2
		} else {
			return 3
		}
	}
	for i := 0; i < n; i++ {
		cnt[charIdx(s[i])] ++
	}
	if check() {
		return 0
	}
	result := n
	for left, right := 0, 0; left < n; left++ {
		for right < n && !check() {
			cnt[charIdx(s[right])] --
			right ++
		}
		if !check() {
			break
		}
		if result > right - left {
			result = right - left
		}
		cnt[charIdx(s[left])] ++
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

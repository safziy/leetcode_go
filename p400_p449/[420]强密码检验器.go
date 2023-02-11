package main

import "unicode"

//如果一个密码满足下述所有条件，则认为这个密码是强密码：
//
//
// 由至少 6 个，至多 20 个字符组成。
// 至少包含 一个小写 字母，一个大写 字母，和 一个数字 。
// 同一字符 不能 连续出现三次 (比如 "...aaa..." 是不允许的, 但是 "...aa...a..." 如果满足其他条件也可以算是强密码)。
//
//
// 给你一个字符串 password ，返回 将 password 修改到满足强密码条件需要的最少修改步数。如果 password 已经是强密码，则返回 0
//。
//
// 在一步修改操作中，你可以：
//
//
// 插入一个字符到 password ，
// 从 password 中删除一个字符，或
// 用另一个字符来替换 password 中的某个字符。
//
//
//
//
// 示例 1：
//
//
//输入：password = "a"
//输出：5
//
//
// 示例 2：
//
//
//输入：password = "aA1"
//输出：3
//
//
// 示例 3：
//
//
//输入：password = "1337C0d3"
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= password.length <= 50
// password 由字母、数字、点 '.' 或者感叹号 '!'
//
//
// Related Topics 贪心 字符串 堆（优先队列） 👍 202 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func strongPasswordChecker(password string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	n, hasLower, hasUpper, hasDigit := len(password), 0, 0, 0
	for _, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = 1
		} else if unicode.IsUpper(ch) {
			hasUpper = 1
		} else if unicode.IsDigit(ch) {
			hasDigit = 1
		}
	}
	typeCnt := hasLower + hasUpper + hasDigit
	if n < 6 {
		return max(6-n, 3-typeCnt)
	} else if n <= 20 {
		replace, cnt, cur := 0, 0, '#'
		for _, ch := range password {
			if ch == cur {
				cnt++
			} else {
				replace += cnt / 3
				cnt = 1
				cur = ch
			}
		}
		replace += cnt / 3
		return max(replace, 3-typeCnt)
	} else {
		// 替换次数和删除次数
		replace, remove := 0, n-20
		// k mod 3 = 1 的组数，即删除 2 个字符可以减少 1 次替换操作
		rm2, cnt, cur := 0, 0, '#'
		for _, ch := range password {
			if ch == cur {
				cnt++
				continue
			}
			if remove > 0 && cnt >= 3 {
				if cnt%3 == 0 {
					// 如果是 k % 3 = 0 的组，那么优先删除 1 个字符，减少 1 次替换操作
					remove--
					replace--
				} else if cnt%3 == 1 {
					// 如果是 k % 3 = 1 的组，那么存下来备用
					rm2++
				}
				//k % 3 = 2 的组无需显式考虑
			}
			replace += cnt / 3
			cnt = 1
			cur = ch
		}
		if remove > 0 && cnt >= 3 {
			if cnt%3 == 0 {
				remove--
				replace--
			} else if cnt%3 == 1 {
				rm2++
			}
		}
		replace += cnt / 3
		// 使用 k % 3 = 1 的组的数量，由剩余的替换次数、组数和剩余的删除次数共同决定
		use2 := min(min(replace, rm2), remove/2)
		replace -= use2
		remove -= use2 * 2
		// 由于每有一次替换次数就一定有 3 个连续相同的字符（k / 3 决定），因此这里可以直接计算出使用 k % 3 = 2 的组的数量
		use3 := min(replace, remove/3)
		replace -= use3
		remove -= use3 * 3
		return (n - 20) + max(replace, 3-typeCnt)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

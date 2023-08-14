package main

import "strings"

//给定一个二进制字符串 s 和一个正整数 n，如果对于 [1, n] 范围内的每个整数，其二进制表示都是 s 的 子字符串 ，就返回 true，否则返回
//false 。
//
// 子字符串 是字符串中连续的字符序列。
//
//
//
// 示例 1：
//
//
//输入：s = "0110", n = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "0110", n = 4
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s[i] 不是 '0' 就是 '1'
// 1 <= n <= 10⁹
//
//
// 👍 125 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func queryString(s string, n int) bool {
	// help函数判断s是否包含[begin, end]的所有二进制字符串，[begin, end]二进制的长度为k
	help := func(k, begin, end int) bool {
		checkMap := make(map[int]bool)
		num := 0
		for i, ch := range s {
			num = num<<1 + int(ch-'0')
			if i >= k {
				num -= int(s[i-k]-'0') << k
			}
			if i >= k-1 && num >= begin && num <= end {
				checkMap[num] = true
			}
		}
		return len(checkMap) == end-begin+1
	}
	// n为1时，只需判断s是否包含"1"
	if n == 1 {
		return strings.Contains(s, "1")
	}
	k := 30
	for (1 << k) >= n {
		k--
	}
	if len(s) < (1<<(k-1))+k-1 || len(s) < n-(1<<k)+k+1 {
		return false
	}
	return help(k, 1<<(k-1), (1<<k)-1) && help(k+1, 1<<k, n)
}

//leetcode submit region end(Prohibit modification and deletion)

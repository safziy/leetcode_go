package main

import "bytes"

//给定一个字符串 s ，根据字符出现的 频率 对其进行 降序排序 。一个字符出现的 频率 是它出现在字符串中的次数。
//
// 返回 已排序的字符串 。如果有多个答案，返回其中任何一个。
//
//
//
// 示例 1:
//
//
//输入: s = "tree"
//输出: "eert"
//解释: 'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//
//
// 示例 2:
//
//
//输入: s = "cccaaa"
//输出: "cccaaa"
//解释: 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//
//
// 示例 3:
//
//
//输入: s = "Aabb"
//输出: "bbAa"
//解释: 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
//
//
//
//
// 提示:
//
//
// 1 <= s.length <= 5 * 10⁵
// s 由大小写英文字母和数字组成
//
//
// 👍 470 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(s string) string {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cntMap := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cntMap[s[i]]++
		maxFreq = max(maxFreq, cntMap[s[i]])
	}
	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cntMap {
		buckets[c] = append(buckets[c], ch)
	}
	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

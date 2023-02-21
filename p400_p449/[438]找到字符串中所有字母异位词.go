package main
//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。 
//
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
// 
//
// 示例 2: 
//
// 
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length, p.length <= 3 * 10⁴ 
// s 和 p 仅包含小写字母 
// 
//
// 👍 1095 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	cnt, sLen, pLen := [26]int{}, len(s), len(p)
	if sLen < pLen {return []int{}}
	for i := 0; i < pLen; i++ {
		cnt[p[i] - 'a'] ++
		cnt[s[i] - 'a'] --
	}
	diff := 0
	for _, c := range cnt {
		if c != 0 {
			diff ++
		}
	}
	result := []int{}
	if diff == 0 {
		result = append(result, 0)
	}
	for start := pLen; start < sLen; start ++ {
		add, sub := s[start] - 'a', s[start-pLen] - 'a'
		if cnt[add] != 0 {
			diff --
		}
		cnt[add] --
		if cnt[add] != 0 {
			diff ++
		}
		if cnt[sub] != 0 {
			diff --
		}
		cnt[sub] ++
		if cnt[sub] != 0 {
			diff ++
		}
		if diff == 0 {
			result = append(result, start-pLen+1)
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

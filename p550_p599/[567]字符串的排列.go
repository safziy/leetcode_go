package main

//给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
//
//
// 示例 2：
//
//
//输入：s1= "ab" s2 = "eidboaoo"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
//
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 855 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	idx := func(ch uint8) int {
		return int(ch - 'a')
	}
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}
	// cnt[i] 表示26个字母的数量的差异 diff表示差异的字母数量
	cnt, diff := [26]int{}, 0
	for i := 0; i < n1; i++ {
		cnt[idx(s1[i])]++
		cnt[idx(s2[i])]--
	}
	for _, c := range cnt {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for e := n1; e < n2; e++ {
		x, y := idx(s2[e-n1]), idx(s2[e])
		if x == y {
			continue
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

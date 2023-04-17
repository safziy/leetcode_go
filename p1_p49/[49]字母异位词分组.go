package main
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。 
//
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。 
//
// 
//
// 示例 1: 
//
// 
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]] 
//
// 示例 2: 
//
// 
//输入: strs = [""]
//输出: [[""]]
// 
//
// 示例 3: 
//
// 
//输入: strs = ["a"]
//输出: [["a"]] 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 10⁴ 
// 0 <= strs[i].length <= 100 
// strs[i] 仅包含小写字母 
// 
//
// 👍 1440 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	charCntMap := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a'] ++
		}
		charCntMap[cnt] = append(charCntMap[cnt], str)
	}
	result := make([][]string, 0, len(charCntMap))
	for _, res := range charCntMap {
		result = append(result, res)
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

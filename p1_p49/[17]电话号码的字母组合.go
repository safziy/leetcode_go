package main

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// 👍 2327 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	mapping := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result, res := []string{}, make([]byte, n)
	var recursion func(int)
	recursion = func(idx int) {
		if idx == n {
			result = append(result, string(res))
			return
		}
		for _, c := range mapping[digits[idx] - '2'] {
			res[idx] = byte(c)
			recursion(idx+1)
		}
	}
	recursion(0)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

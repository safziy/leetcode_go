package main
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：["()"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 8 
// 
//
// 👍 3103 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	res, result := make([]byte, 0, n << 1), []string{}
	var recursion func(left, right int)
	recursion = func(left, right int) {
		if left == n && right == n {
			result = append(result, string(res))
			return
		}
		if left < n {
			res = append(res, '(')
			recursion(left+1, right)
			res = res[:len(res)-1]
		}
		if right < left {
			res = append(res, ')')
			recursion(left, right+1)
			res = res[:len(res)-1]
		}
	}
	recursion(0, 0)
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

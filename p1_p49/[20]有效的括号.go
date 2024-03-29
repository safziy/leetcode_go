package main
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 每个右括号都有一个对应的相同类型的左括号。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "()[]{}"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：s = "(]"
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁴ 
// s 仅由括号 '()[]{}' 组成 
// 
//
// 👍 3826 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := []byte{}
	for _, ch := range s {
		if ch == ')' {
			if len(stack) <= 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if ch == ']' {
			if len(stack) <= 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if ch == '}' {
			if len(stack) <= 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(ch))
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

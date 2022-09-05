package main

import "fmt"

//求解一个给定的方程，将x以字符串 "x=#value" 的形式返回。该方程仅包含 '+' ， '-' 操作，变量 x 和其对应系数。
//
// 如果方程没有解或存在的解不为整数，请返回 "No solution" 。如果方程有无限解，则返回 “Infinite solutions” 。 
//
// 题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。 
//
// 
//
// 示例 1： 
//
// 
//输入: equation = "x+5-3+x=6+x-2"
//输出: "x=2"
// 
//
// 示例 2: 
//
// 
//输入: equation = "x=x"
//输出: "Infinite solutions"
// 
//
// 示例 3: 
//
// 
//输入: equation = "2x=x"
//输出: "x=0"
// 
//
// 
//
// 提示: 
//
// 
// 3 <= equation.length <= 1000 
// equation 只有一个 '='. 
// 方程由绝对值在 [0, 100] 范围内且无任何前导零的整数和变量 'x' 组成。 
// 
//
// Related Topics 数学 字符串 模拟 👍 193 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func solveEquation(equation string) string {
	a, b, t, s1, s2 := 0, 0, 0, 1, 1
	for _, item := range equation {
		switch item {
		case '+':
			b += s1 * s2 * t
			s1 = 1
			t = 0
		case '-':
			b += s1 * s2 * t
			s1 = -1
			t = 0
		case '=':
			b += s1 * s2 * t
			s1 = 1
			s2 = -1
			t = 0
		case 'x':
			if t == 0 {
				a += s1 * s2
			} else {
				a += s1 * s2 * t
				t = 0
			}
		default:
			t = t*10 + int(item-'0')
		}
	}
	if t != 0 {
		b += s1 * s2 * t
	}
	if a == 0 {
		if b == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		return fmt.Sprintf("x=%d", -b/a)
	}
}
//leetcode submit region end(Prohibit modification and deletion)

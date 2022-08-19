package p600_p649

import (
	"fmt"
)

/*
640. 求解方程

求解一个给定的方程，将x以字符串 "x=#value"的形式返回。该方程仅包含 '+' ， '-' 操作，变量x和其对应系数。

如果方程没有解，请返回"No solution"。如果方程有无限解，则返回 “Infinite solutions” 。

题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。


提示：

3 <= equation.length <= 1000
equation只有一个'='.
equation方程由整数组成，其绝对值在[0, 100]范围内，不含前导零和变量 'x' 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/solve-the-equation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
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
	fmt.Println(a, b, t)
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

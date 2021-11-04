package p350_p399

/*
367. 有效的完全平方数
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如  sqrt 。

提示：

1 <= num <= 2^31 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-perfect-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPerfectSquare(num int) bool {
	// 二分查找
	//min, max := 1, num
	//for min <= max {
	//	mid := (min + max) >> 1
	//	square := mid * mid
	//	if square == num {
	//		return true
	//	} else if square > num {
	//		max = mid - 1
	//	} else if square < num {
	//		min = mid + 1
	//	}
	//}
	//return false
	// 牛顿迭代法
	x0 := float64(num)
	for {
		x1 := (x0 + float64(num)/x0) / 2
		if x0-x1 < 1e-6 {
			break
		}
		x0 = x1
	}
	return int(x0)*int(x0) == num
}

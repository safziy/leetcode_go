package main
//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。 
//
// 进阶：不要 使用任何内置的库函数，如 sqrt 。 
//
// 
//
// 示例 1： 
//
// 
//输入：num = 16
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：num = 14
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= num <= 2^31 - 1 
// 
//
// Related Topics 数学 二分查找 👍 432 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
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
//leetcode submit region end(Prohibit modification and deletion)

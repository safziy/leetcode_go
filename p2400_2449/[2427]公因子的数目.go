package main

//给你两个正整数 a 和 b ，返回 a 和 b 的 公 因子的数目。
//
// 如果 x 可以同时整除 a 和 b ，则认为 x 是 a 和 b 的一个 公因子 。
//
//
//
// 示例 1：
//
// 输入：a = 12, b = 6
//输出：4
//解释：12 和 6 的公因子是 1、2、3、6 。
//
//
// 示例 2：
//
// 输入：a = 25, b = 30
//输出：2
//解释：25 和 30 的公因子是 1、5 。
//
//
//
// 提示：
//
//
// 1 <= a, b <= 1000
//
//
// 👍 44 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func commonFactors(a int, b int) int {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	c, result := gcd(a, b), 0
	for i := 1; i*i <= c; i++ {
		if c%i == 0 {
			result++
			if i*i != c {
				result++
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

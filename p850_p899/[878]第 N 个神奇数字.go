package main

//一个正整数如果能被 a 或 b 整除，那么它是神奇的。
//
// 给定三个整数 n , a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案 对 10⁹ + 7 取模 后的值。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：n = 1, a = 2, b = 3
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：n = 4, a = 2, b = 3
//输出：6
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 10⁹ 
// 2 <= a, b <= 4 * 10⁴ 
// 
//
// 
//
// 👍 214 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func nthMagicalNumber(n int, a int, b int) int {
	const mod int = 1e9 + 7
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	l, r, c := min(a, b), n*min(a, b), a/gcd(a, b)*b
	for l <= r {
		mid := (l + r) >> 1
		cnt := mid/a + mid/b - mid/c
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod
}

//leetcode submit region end(Prohibit modification and deletion)

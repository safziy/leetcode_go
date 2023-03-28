package main

//给你四个整数：n 、a 、b 、c ，请你设计一个算法来找出第 n 个丑数。
//
// 丑数是可以被 a 或 b 或 c 整除的 正整数 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, a = 2, b = 3, c = 5
//输出：4
//解释：丑数序列为 2, 3, 4, 5, 6, 8, 9, 10... 其中第 3 个是 4。
//
// 示例 2：
//
//
//输入：n = 4, a = 2, b = 3, c = 4
//输出：6
//解释：丑数序列为 2, 3, 4, 6, 8, 9, 10, 12... 其中第 4 个是 6。
//
//
// 示例 3：
//
//
//输入：n = 5, a = 2, b = 11, c = 13
//输出：10
//解释：丑数序列为 2, 4, 6, 8, 10, 11, 12, 13... 其中第 5 个是 10。
//
//
// 示例 4：
//
//
//输入：n = 1000000000, a = 2, b = 217983653, c = 336916467
//输出：1999999984
//
//
//
//
// 提示：
//
//
// 1 <= n, a, b, c <= 10^9
// 1 <= a * b * c <= 10^18
// 本题结果在 [1, 2 * 10^9] 的范围内
//
//
// 👍 126 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int, a int, b int, c int) int {
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
	l := min(a, min(b, c))
	r, ab, bc, ac := n*l, a/gcd(a, b)*b, b/gcd(b, c)*c, a/gcd(a, c)*c
	abc := ab / gcd(ab, c) * c
	for l <= r {
		mid := (l + r) >> 1
		cnt := mid/a + mid/b + mid/c - mid/ab - mid/bc - mid/ac + mid/abc
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r + 1
}

//leetcode submit region end(Prohibit modification and deletion)

package main

//给定正整数 k ，你需要找出可以被 k 整除的、仅包含数字 1 的最 小 正整数 n 的长度。
//
// 返回 n 的长度。如果不存在这样的 n ，就返回-1。
//
// 注意： n 不符合 64 位带符号整数。
//
//
//
// 示例 1：
//
//
//输入：k = 1
//输出：1
//解释：最小的答案是 n = 1，其长度为 1。
//
// 示例 2：
//
//
//输入：k = 2
//输出：-1
//解释：不存在可被 2 整除的正整数 n 。
//
// 示例 3：
//
//
//输入：k = 3
//输出：3
//解释：最小的答案是 n = 111，其长度为 3。
//
//
//
// 提示：
//
//
// 1 <= k <= 10⁵
//
//
// 👍 134 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	// cur 代表当前的余数，初始值为 0，res 代表数字 1 的个数，初始值为 1。
	cur, result := 0, 1
	// 不断计算下一个数字的余数，直到找到一个可行的解或者发现不存在可行的解。
	for {
		cur = (10*cur + 1) % k
		if cur == 0 {
			return result
		}
		result++
	}
}

//leetcode submit region end(Prohibit modification and deletion)

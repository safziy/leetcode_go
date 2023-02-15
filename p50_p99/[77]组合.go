package main

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//
// 示例 2：
//
//
//输入：n = 1, k = 1
//输出：[[1]]
//
//
//
// 提示：
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// Related Topics 回溯 👍 1278 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	cnt := 1
	for i := 1; i <= k; i++ {
		cnt *= (n - k + 1) / i
	}
	result := make([][]int, 0, cnt)
	var recursion func(cur []int, start, count int)
	recursion = func(cur []int, start, count int) {
		if count == 0 {
			comb := make([]int, 0, k)
			copy(comb, cur)
			result = append(result, comb)
			return
		}
		end, pos := n-count+1, k-count
		for i := start; i <= end; i++ {
			cur[pos] = i
			recursion(cur, i+1, count-1)
		}
	}
	cur := make([]int, k)
	recursion(cur, 1, k)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

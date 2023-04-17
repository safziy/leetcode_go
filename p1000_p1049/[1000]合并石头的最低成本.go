package main

//有 N 堆石头排成一排，第 i 堆中有 stones[i] 块石头。
//
// 每次移动（move）需要将连续的 K 堆石头合并为一堆，而这个移动的成本为这 K 堆石头的总数。
//
// 找出把所有石头合并成一堆的最低成本。如果不可能，返回 -1 。
//
//
//
// 示例 1：
//
// 输入：stones = [3,2,4,1], K = 2
//输出：20
//解释：
//从 [3, 2, 4, 1] 开始。
//合并 [3, 2]，成本为 5，剩下 [5, 4, 1]。
//合并 [4, 1]，成本为 5，剩下 [5, 5]。
//合并 [5, 5]，成本为 10，剩下 [10]。
//总成本 20，这是可能的最小值。
//
//
// 示例 2：
//
// 输入：stones = [3,2,4,1], K = 3
//输出：-1
//解释：任何合并操作后，都会剩下 2 堆，我们无法再进行合并。所以这项任务是不可能完成的。.
//
//
// 示例 3：
//
// 输入：stones = [3,5,1,2,6], K = 3
//输出：25
//解释：
//从 [3, 5, 1, 2, 6] 开始。
//合并 [5, 1, 2]，成本为 8，剩下 [3, 8, 6]。
//合并 [3, 8, 6]，成本为 17，剩下 [17]。
//总成本 25，这是可能的最小值。
//
//
//
//
// 提示：
//
//
// 1 <= stones.length <= 30
// 2 <= K <= 30
// 1 <= stones[i] <= 100
//
//
// 👍 249 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func mergeStones(stones []int, k int) int {
	const INF = 0x3f3f3f3f
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	// 方法一: 动态规划+记忆化搜索

	// memo[key]表示[s,e]合并为t堆的最优解 key = (s * 32 + e) * 32 + t
	memo := make(map[int]int, n*n)
	// 前缀和减少求和运算
	preSum := make([]int, n+1)
	for i, stone := range stones {
		preSum[i+1] = preSum[i] + stone
		memo[(i<<5 +i) << 5 + 1] = 0
	}
	var search func(l, r, t int) int
	search = func(l, r, t int) int {
		key := (l<<5 +r) << 5 + t
		if res, ok := memo[key]; ok {
			return res
		}
		if t > r-l+1 {
			return INF
		}
		if t == 1 {
			res := search(l, r, k)
			if res == INF {
				memo[key] = INF
			} else {
				memo[key] = res + (preSum[r+1] - preSum[l])
			}
		} else {
			res := INF
			for i := l; i < r; i += k-1 {
				res = min(res, search(l, i, 1) + search(i + 1, r, t - 1))
			}
			memo[key] = res
		}
		return memo[key]
	}
	return search(0, n-1, 1)
}

//leetcode submit region end(Prohibit modification and deletion)

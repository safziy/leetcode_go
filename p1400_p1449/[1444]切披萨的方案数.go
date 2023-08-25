package main

//给你一个 rows x cols 大小的矩形披萨和一个整数 k ，矩形包含两种字符： 'A' （表示苹果）和 '.' （表示空白格子）。你需要切披萨 k-1
// 次，得到 k 块披萨并送给别人。
//
// 切披萨的每一刀，先要选择是向垂直还是水平方向切，再在矩形的边界上选一个切的位置，将披萨一分为二。如果垂直地切披萨，那么需要把左边的部分送给一个人，如果水平
//地切，那么需要把上面的部分送给一个人。在切完最后一刀后，需要把剩下来的一块送给最后一个人。
//
// 请你返回确保每一块披萨包含 至少 一个苹果的切披萨方案数。由于答案可能是个很大的数字，请你返回它对 10^9 + 7 取余的结果。
//
//
//
// 示例 1：
//
//
//
// 输入：pizza = ["A..","AAA","..."], k = 3
//输出：3
//解释：上图展示了三种切披萨的方案。注意每一块披萨都至少包含一个苹果。
//
//
// 示例 2：
//
// 输入：pizza = ["A..","AA.","..."], k = 3
//输出：1
//
//
// 示例 3：
//
// 输入：pizza = ["A..","A..","..."], k = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= rows, cols <= 50
// rows == pizza.length
// cols == pizza[i].length
// 1 <= k <= 10
// pizza 只包含字符 'A' 和 '.' 。
//
//
// 👍 142 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func ways(pizza []string, k int) int {
	var MOD int = 1e9 + 7
	rows, cols := len(pizza), len(pizza[0])
	// 统计一下每个位置右下方有多少个苹果
	cnt := make([][]int, rows+1)
	for i := range cnt {
		cnt[i] = make([]int, cols+1)
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			cnt[i][j] = cnt[i+1][j] + cnt[i][j+1] - cnt[i+1][j+1]
			if pizza[i][j] == 'A' {
				cnt[i][j]++
			}
		}
	}
	if cnt[0][0] < k {
		return 0
	}
	// 记忆化搜索
	cache := make(map[int]int)
	var search func(r, c, l int) int
	search = func(r, c, l int) int {
		key := r<<20 + c<<10 + l
		if val, ok := cache[key]; ok {
			return val
		}
		if cnt[r][c] == 0 {
			cache[key] = 0
			return 0
		}
		if l == 1 {
			cache[key] = 1
			return 1
		}
		var ans int
		for i := r; i < rows-1; i++ {
			if cnt[r][c]-cnt[i+1][c] > 0 {
				ans += search(i+1, c, l-1)
			}
		}
		for j := c; j < cols-1; j++ {
			if cnt[r][c]-cnt[r][j+1] > 0 {
				ans += search(r, j+1, l-1)
			}
		}
		ans %= MOD
		cache[key] = ans
		return ans
	}
	return search(0, 0, k)
}

//leetcode submit region end(Prohibit modification and deletion)

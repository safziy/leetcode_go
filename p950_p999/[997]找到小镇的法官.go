package main

//小镇里有 n 个人，按从 1 到 n 的顺序编号。传言称，这些人中有一个暗地里是小镇法官。
//
// 如果小镇法官真的存在，那么： 
//
// 
// 小镇法官不会信任任何人。 
// 每个人（除了小镇法官）都信任这位小镇法官。 
// 只有一个人同时满足属性 1 和属性 2 。 
// 
//
// 给你一个数组 trust ，其中 trust[i] = [ai, bi] 表示编号为 ai 的人信任编号为 bi 的人。 
//
// 如果小镇法官存在并且可以确定他的身份，请返回该法官的编号；否则，返回 -1 。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 2, trust = [[1,2]]
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：n = 3, trust = [[1,3],[2,3]]
//输出：3
// 
//
// 示例 3： 
//
// 
//输入：n = 3, trust = [[1,3],[2,3],[3,1]]
//输出：-1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 1000 
// 0 <= trust.length <= 10⁴ 
// trust[i].length == 2 
// trust 中的所有trust[i] = [ai, bi] 互不相同 
// ai != bi 
// 1 <= ai, bi <= n 
// 
//
// 👍 290 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findJudge(n int, trust [][]int) int {
	// in 表示有多少人信任他 out 表示他信任多少人
	in, out := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		in[t[1]]++
		out[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

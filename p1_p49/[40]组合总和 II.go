package main

import "sort"

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。 
//
// 注意：解集不能包含重复的组合。 
//
// 
//
// 示例 1: 
//
// 
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//] 
//
// 示例 2: 
//
// 
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//] 
//
// 
//
// 提示: 
//
// 
// 1 <= candidates.length <= 100 
// 1 <= candidates[i] <= 50 
// 1 <= target <= 30 
// 
//
// 👍 1248 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	n, result := len(candidates), [][]int{}
	sort.Ints(candidates)
	res, sum := make([]int, 0, n), 0
	var recursion func(int)
	recursion = func(idx int) {
		for i := idx; i < n; i++ {
			sum += candidates[i]
			res = append(res, candidates[i])
			if sum >= target {
				if sum == target {
					result = append(result, append([]int{}, res...))
				}
				res = res[:len(res)-1]
				sum -= candidates[i]
				break
			}
			recursion(i+1)
			res = res[:len(res)-1]
			sum -= candidates[i]
			for i < n - 1 && candidates[i] == candidates[i+1] {
				i++
			}
		}
	}
	recursion(0)
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

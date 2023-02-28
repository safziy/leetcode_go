package main

import "sort"

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。 
//
// 
// 
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0]
//输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// 
//
// 👍 1031 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n, res, result := len(nums), []int{}, [][]int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, idx int) {
		if idx == n {
			result = append(result, append([]int{}, res...))
			return
		}
		// 不选择当前的数字
		dfs(false, idx+1)
		if !choosePre && idx > 0 && nums[idx-1] == nums[idx] {
			// 相同的数字处理
			return
		}
		// 选择当前的数字
		res = append(res, nums[idx])
		dfs(true, idx+1)
		res = res[:len(res)-1]
	}
	dfs(false, 0)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

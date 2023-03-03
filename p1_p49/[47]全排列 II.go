package main

import "sort"

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// 👍 1296 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	n, result := len(nums), [][]int{}
	sort.Ints(nums)
	flag, res := make([]bool, n), make([]int, 0, n)
	var recursion func()
	recursion = func() {
		if len(res) == n {
			result = append(result, append([]int{}, res...))
			return
		}
		for i := 0; i < n; i++ {
			if flag[i] {
				continue
			}
			flag[i] = true
			res = append(res, nums[i])
			recursion()
			res = res[:len(res)-1]
			flag[i] = false
			for i < n-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	recursion()
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

package main
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 6 
// -10 <= nums[i] <= 10 
// nums 中的所有整数 互不相同 
// 
//
// Related Topics 数组 回溯 👍 2401 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	n := len(nums)
	cnt := 1
	for i := 1; i <= n; i++ {
		cnt *= i
	}
	result := make([][]int, 0, cnt)
	var recursion func(comb []int)
	recursion = func(comb []int) {
		l := len(comb)
		if l == n {
			result = append(result, comb)
			return
		}
		for i := 0; i <= l; i++ {
			temp := make([]int, 0, n)
			temp = append(temp, comb[0:i]...)
			temp = append(temp, nums[l])
			temp = append(temp, comb[i:l]...)
			recursion(temp)
		}
	}
	recursion(make([]int, 0))
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

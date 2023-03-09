package main

//给你一个整数数组 nums ，返回其中 按位与三元组 的数目。
//
// 按位与三元组 是由下标 (i, j, k) 组成的三元组，并满足下述全部条件： 
//
// 
// 0 <= i < nums.length 
// 0 <= j < nums.length 
// 0 <= k < nums.length 
// nums[i] & nums[j] & nums[k] == 0 ，其中 & 表示按位与运算符。 
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,1,3]
//输出：12
//解释：可以选出如下 i, j, k 三元组：
//(i=0, j=0, k=1) : 2 & 2 & 1
//(i=0, j=1, k=0) : 2 & 1 & 2
//(i=0, j=1, k=1) : 2 & 1 & 1
//(i=0, j=1, k=2) : 2 & 1 & 3
//(i=0, j=2, k=1) : 2 & 3 & 1
//(i=1, j=0, k=0) : 1 & 2 & 2
//(i=1, j=0, k=1) : 1 & 2 & 1
//(i=1, j=0, k=2) : 1 & 2 & 3
//(i=1, j=1, k=0) : 1 & 1 & 2
//(i=1, j=2, k=0) : 1 & 3 & 2
//(i=2, j=0, k=1) : 3 & 2 & 1
//(i=2, j=1, k=0) : 3 & 1 & 2
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,0,0]
//输出：27
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 1000 
// 0 <= nums[i] < 2¹⁶ 
// 
//
// 👍 117 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func countTriplets(nums []int) int {
	cnt := make([]int, 1<<16)
	for _, num1 := range nums {
		for _, num2 := range nums {
			cnt[num1&num2]++
		}
	}
	result := 0
	for _, num := range nums {
		x := num ^ 0xffff
		for sub := x; sub > 0; sub = (sub - 1) & x {
			result += cnt[sub]
		}
		result += cnt[0]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

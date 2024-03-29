package main
//给你一个下标从 0 开始、严格递增 的整数数组 nums 和一个正整数 diff 。如果满足下述全部条件，则三元组 (i, j, k) 就是一个 算术三元组
// ： 
//
// 
// i < j < k ， 
// nums[j] - nums[i] == diff 且 
// nums[k] - nums[j] == diff 
// 
//
// 返回不同 算术三元组 的数目。 
//
// 
//
// 示例 1： 
//
// 输入：nums = [0,1,4,6,7,10], diff = 3
//输出：2
//解释：
//(1, 2, 4) 是算术三元组：7 - 4 == 3 且 4 - 1 == 3 。
//(2, 4, 5) 是算术三元组：10 - 7 == 3 且 7 - 4 == 3 。
// 
//
// 示例 2： 
//
// 输入：nums = [4,5,6,7,8,9], diff = 2
//输出：2
//解释：
//(0, 2, 4) 是算术三元组：8 - 6 == 2 且 6 - 4 == 2 。
//(1, 3, 5) 是算术三元组：9 - 7 == 2 且 7 - 5 == 2 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 200 
// 0 <= nums[i] <= 200 
// 1 <= diff <= 50 
// nums 严格 递增 
// 
//
// 👍 46 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func arithmeticTriplets(nums []int, diff int) int {
	// 方法一: 三指针
	n, result := len(nums), 0
	for i, j, k := 0, 1, 2; i < n-2 && j < n-1 && k < n; i++ {
		for j < n -1 && nums[j] - nums[i] < diff {
			j++
		}
		if j >= n-1 || nums[j] - nums[i] > diff {
			continue
		}
		for k < n && nums[k] - nums[j] < diff {
			k++
		}
		if k < n && nums[k]-nums[j] == diff {
			result ++
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

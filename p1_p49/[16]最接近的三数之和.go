package main

import "sort"

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。 
//
// 假定每组输入只存在恰好一个解。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,0,0], target = 1
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 1000 
// -1000 <= nums[i] <= 1000 
// -10⁴ <= target <= 10⁴ 
// 
//
// 👍 1450 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left, right := i + 1, len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum - target) < abs(result - target) {
				result = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return result
			}
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

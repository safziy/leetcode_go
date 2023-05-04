package main

//给你一个整数数组 nums 和两个整数 firstLen 和 secondLen，请你找出并返回两个非重叠 子数组 中元素的最大和，长度分别为
//firstLen 和 secondLen 。
//
// 长度为 firstLen 的子数组可以出现在长为 secondLen 的子数组之前或之后，但二者必须是不重叠的。
//
// 子数组是数组的一个 连续 部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
//输出：20
//解释：子数组的一种选择中，[9] 长度为 1，[6,5] 长度为 2。
//
//
// 示例 2：
//
//
//输入：nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
//输出：29
//解释：子数组的一种选择中，[3,8,1] 长度为 3，[8,9] 长度为 2。
//
//
// 示例 3：
//
//
//输入：nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
//输出：31
//解释：子数组的一种选择中，[5,6,0,9] 长度为 4，[0,3,8] 长度为 3。
//
//
//
//
// 提示：
//
//
// 1 <= firstLen, secondLen <= 1000
// 2 <= firstLen + secondLen <= 1000
// firstLen + secondLen <= nums.length <= 1000
// 0 <= nums[i] <= 1000
//
//
// 👍 220 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(nums)
	preMax, sufMax := make([]int, n), make([]int, n)
	sum, tempMax := 0, 0
	for i := 0; i < n; i++ {
		if i >= firstLen {
			sum -= nums[i-firstLen]
		}
		sum += nums[i]
		if i >= firstLen-1 {
			tempMax = max(tempMax, sum)
			preMax[i] = tempMax
		}
	}
	sum, tempMax = 0, 0
	for i := n - 1; i >= 0; i-- {
		if i < n-firstLen {
			sum -= nums[i+firstLen]
		}
		sum += nums[i]
		if i <= n-firstLen {
			tempMax = max(tempMax, sum)
			sufMax[i] = tempMax
		}
	}
	result := 0
	sum = 0
	for i := 0; i < n; i++ {
		if i >= secondLen {
			sum -= nums[i-secondLen]
		}
		sum += nums[i]
		if i >= secondLen-1 {
			if i-secondLen >= firstLen-1 {
				result = max(result, sum+preMax[i-secondLen])
			}
			if n-i-1 >= firstLen {
				result = max(result, sum+sufMax[i+1])
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

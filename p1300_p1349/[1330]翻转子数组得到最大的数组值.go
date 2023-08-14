package p1300_p1349

import "math"

//给你一个整数数组 nums 。「数组值」定义为所有满足 0 <= i < nums.length-1 的 |nums[i]-nums[i+1]| 的和。
//
// 你可以选择给定数组的任意子数组，并将该子数组翻转。但你只能执行这个操作 一次 。
//
// 请你找到可行的最大 数组值 。
//
//
//
// 示例 1：
//
// 输入：nums = [2,3,1,5,4]
//输出：10
//解释：通过翻转子数组 [3,1,5] ，数组变成 [2,5,1,3,4] ，数组值为 10 。
//
//
// 示例 2：
//
// 输入：nums = [2,4,9,24,2,1,10]
//输出：68
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3*10^4
// -10^5 <= nums[i] <= 10^5
//
//
// 👍 132 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxValueAfterReverse(nums []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	base, d, n := 0, 0, len(nums)
	mx, mn := math.MinInt, math.MaxInt
	for i := 1; i < n; i++ {
		a, b := nums[i-1], nums[i]
		base += abs(a - b)
		mx = max(mx, min(a, b))
		mn = min(mn, max(a, b))
		d = max(d, max(abs(nums[0]-b)-abs(a-b), // i=0
			abs(nums[n-1]-a)-abs(a-b))) // j=n-1
	}
	return base + max(d, 2*(mx-mn))
}

//leetcode submit region end(Prohibit modification and deletion)

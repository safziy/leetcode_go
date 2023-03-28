package main

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// 👍 4208 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, result := len(height), 0
	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if height[left] < height[right] {
			result += leftMax - height[left]
			left++
			leftMax = max(leftMax, height[left])
		} else {
			result += rightMax - height[right]
			right--
			rightMax = max(rightMax, height[right])
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

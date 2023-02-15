package main

//给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
//
// 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
//
// 所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
//
// 请你返回「表现良好时间段」的最大长度。
//
//
//
// 示例 1：
//
//
//输入：hours = [9,9,6,0,6,6,9]
//输出：3
//解释：最长的表现良好时间段是 [9,9,6]。
//
// 示例 2：
//
//
//输入：hours = [6,6,6]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= hours.length <= 10⁴
// 0 <= hours[i] <= 16
//
//
// Related Topics 栈 数组 哈希表 前缀和 单调栈 👍 292 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestWPI(hours []int) int {
	// 前缀和+单调栈
	//n := len(hours)
	//s := make([]int, n+1)
	//stack, sIdx := make([]int, n+1), 1
	//for i := 1; i <= n; i++ {
	//	if hours[i-1] > 8 {
	//		s[i] = s[i-1] + 1
	//	} else {
	//		s[i] = s[i-1] - 1
	//	}
	//	if s[stack[sIdx-1]] > s[i] {
	//		stack[sIdx] = i
	//		sIdx++
	//	}
	//}
	//result := 0
	//for r := n; r >= 1; r-- {
	//	for sIdx > 0 && s[stack[sIdx-1]] < s[r] {
	//		if r-stack[sIdx-1] > result {
	//			result = r - stack[sIdx-1]
	//		}
	//		sIdx--
	//	}
	//}
	//return result

	// 前缀和+hashmap
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, sum, result := len(hours), 0, 0
	m := map[int]int{}
	for i := 1; i <= n; i++ {
		if hours[i-1] > 8 {
			sum = sum + 1
		} else {
			sum = sum - 1
		}
		if sum > 0 {
			result = max(result, i)
		} else {
			if m[sum-1] != 0 {
				result = max(result, i-m[sum-1])
			}
			if m[sum] == 0 {
				m[sum] = i
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

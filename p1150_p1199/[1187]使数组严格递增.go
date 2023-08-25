package p1150_p1199

import (
	"math"
	"sort"
)

//给你两个整数数组 arr1 和 arr2，返回使 arr1 严格递增所需要的最小「操作」数（可能为 0）。
//
// 每一步「操作」中，你可以分别从 arr1 和 arr2 中各选出一个索引，分别为 i 和 j，0 <= i < arr1.length 和 0 <= j
//< arr2.length，然后进行赋值运算 arr1[i] = arr2[j]。
//
// 如果无法让 arr1 严格递增，请返回 -1。
//
//
//
// 示例 1：
//
// 输入：arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
//输出：1
//解释：用 2 来替换 5，之后 arr1 = [1, 2, 3, 6, 7]。
//
//
// 示例 2：
//
// 输入：arr1 = [1,5,3,6,7], arr2 = [4,3,1]
//输出：2
//解释：用 3 来替换 5，然后用 4 来替换 3，得到 arr1 = [1, 3, 4, 6, 7]。
//
//
// 示例 3：
//
// 输入：arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
//输出：-1
//解释：无法使 arr1 严格递增。
//
//
//
// 提示：
//
//
// 1 <= arr1.length, arr2.length <= 2000
// 0 <= arr1[i], arr2[i] <= 10^9
//
//
//
//
// 👍 160 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	sort.Ints(arr2)
	n1 := len(arr1)
	n2 := len(arr2)
	// 动态规划一
	// dp[i][j] 表示数组 arr1 中的前 i 个元素进行了 j 次替换后组成严格递增子数组末尾元素的最小值
	dp := make([][]int, n1+1)
	for i := range dp {
		//dp[i] = make([]int, n1+1)
		dp[i] = make([]int, min(n1, n2)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = -1
	for i := 1; i <= n1; i++ {
		for j := 0; j <= min(i, n2); j++ {
			if arr1[i-1] > dp[i-1][j] {
				dp[i][j] = arr1[i-1]
			}
			if j > 0 && dp[i-1][j-1] != math.MaxInt {
				k := j - 1 + sort.SearchInts(arr2[j-1:], dp[i-1][j-1]+1)
				if k < n2 {
					dp[i][j] = min(dp[i][j], arr2[k])
				}
			}
			if i == n1 && dp[i][j] != math.MaxInt {
				return j
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

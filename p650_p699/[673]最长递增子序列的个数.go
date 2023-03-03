package main

import "sort"

//给定一个未排序的整数数组
// nums ， 返回最长递增子序列的个数 。
//
// 注意 这个数列必须是 严格 递增的。
//
//
//
// 示例 1:
//
//
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//
//
// 示例 2:
//
//
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
//
//
//
// 提示:
//
//
//
//
//
// 1 <= nums.length <= 2000
// -10⁶ <= nums[i] <= 10⁶
//
//
// 👍 705 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberOfLIS(nums []int) int {
	d := [][]int{}
	cnt := [][]int{}
	for _, v := range nums {
		i := sort.Search(len(d), func(i int) bool { return d[i][len(d[i])-1] >= v })
		c := 1
		if i > 0 {
			k := sort.Search(len(d[i-1]), func(k int) bool { return d[i-1][k] < v })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}
		if i == len(d) {
			d = append(d, []int{v})
			cnt = append(cnt, []int{0, c})
		} else {
			d[i] = append(d[i], v)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+c)
		}
	}
	c := cnt[len(cnt)-1]
	return c[len(c)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

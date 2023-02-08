package main

//给你一个由 不同 正整数组成的数组 nums ，请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量。其中 a、b、c 和
//d 都是 nums 中的元素，且 a != b != c != d 。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,4,6]
//输出：8
//解释：存在 8 个满足题意的元组：
//(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
//(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
//
//
// 示例 2：
//
//
//输入：nums = [1,2,4,5,10]
//输出：16
//解释：存在 16 个满足题意的元组：
//(1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
//(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
//(2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
//(4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 10⁴
// nums 中的所有元素 互不相同
//
//
// Related Topics 数组 哈希表 👍 24 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func tupleSameProduct(nums []int) int {
	n := len(nums)
	m := make(map[int]int, n)
	cnt := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			mul := nums[i] * nums[j]
			cnt += m[mul]
			m[mul] = m[mul] + 1
		}
	}
	return cnt << 3
}

//leetcode submit region end(Prohibit modification and deletion)

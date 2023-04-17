package p1050_p1099

//给你一个正整数数组 arr（可能存在重复的元素），请你返回可在 一次交换（交换两数字 arr[i] 和 arr[j] 的位置）后得到的、按字典序排列小于
//arr 的最大排列。
//
// 如果无法这么操作，就请返回原数组。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,2,1]
//输出：[3,1,2]
//解释：交换 2 和 1
//
//
// 示例 2：
//
//
//输入：arr = [1,1,5]
//输出：[1,1,5]
//解释：已经是最小排列
//
//
// 示例 3：
//
//
//输入：arr = [1,9,4,6,7]
//输出：[1,7,4,6,9]
//解释：交换 9 和 7
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10⁴
// 1 <= arr[i] <= 10⁴
//
//
// 👍 110 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func prevPermOpt1(arr []int) []int {
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			j := len(arr) - 1
			for arr[j] >= arr[i] || arr[j] == arr[j-1] {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

package p250_p299

/*
268. 丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

提示：

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
nums 中的所有数字都 独一无二

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func missingNumber(nums []int) int {
	// map求解
	//check := map[int]bool{}
	//for _, num := range nums {
	//	check[num] = true
	//}
	//for i := 0; ; i++ {
	//	if !check[i] {
	//		return i
	//	}
	//}
	// 位运算
	//xor := 0
	//for i, num := range nums {
	//	xor ^= i ^ num
	//}
	//return xor ^ len(nums)
	// 求和
	n := len(nums)
	total := n * (n + 1) / 2
	for _, num := range nums {
		total -= num
	}
	return total
}

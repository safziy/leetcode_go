package main

//给你一个数组 nums ，请你完成两类查询。
//
//
// 其中一类查询要求 更新 数组 nums 下标对应的值
// 另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
//
//
// 实现 NumArray 类：
//
//
// NumArray(int[] nums) 用整数数组 nums 初始化对象
// void update(int index, int val) 将 nums[index] 的值 更新 为 val
// int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元
//素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
//
//
//
//
// 示例 1：
//
//
//输入：
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//输出：
//[null, 9, null, 8]
//
//解释：
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
//numArray.update(1, 2);   // nums = [1,2,5]
//numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -100 <= nums[i] <= 100
// 0 <= index < nums.length
// -100 <= val <= 100
// 0 <= left <= right < nums.length
// 调用 update 和 sumRange 方法次数不大于 3 * 10⁴
//
//
// 👍 589 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	src, tree []int
	n int
}

func lowBit(x int) int {
	return x & -x
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1)
	na := NumArray{src: nums, tree: tree, n: n}
	for i := 0; i < n; i++ {
		na.add(i+1, nums[i])
	}
	return na
}

func (this *NumArray) add(index int, val int) {
	for ; index <= this.n; index += index & -index {
		this.tree[index] += val
	}
}

func (this *NumArray) prefixSum(index int) (sum int) {
	for ;index > 0; index -= index & -index {
		sum += this.tree[index]
	}
	return
}

func (this *NumArray) Update(index int, val int) {
	this.add(index + 1, val - this.src[index])
	this.src[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum(right+1) - this.prefixSum(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

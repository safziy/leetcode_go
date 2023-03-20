package main

//给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。
//
//
//
// 示例 1：
//
//
//输入：n = 20
//输出：1
//解释：具有至少 1 位重复数字的正数（<= 20）只有 11 。
//
//
// 示例 2：
//
//
//输入：n = 100
//输出：10
//解释：具有至少 1 位重复数字的正数（<= 100）有 11，22，33，44，55，66，77，88，99 和 100 。
//
//
// 示例 3：
//
//
//输入：n = 1000
//输出：262
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
//
//
// 👍 157 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func numDupDigitsAtMostN(n int) int {
	if n <= 10 {
		return 0
	}
	// 计算 A(b, l) 即从b个数字中带顺序取l个数字
	a := func(b, l int) int {
		res := 1
		for i := 0; i < l; i++ {
			res *= b
			b--
		}
		return res
	}
	// baseVal[i]代表i位数中有多少个不含重复数字的数字个数
	//baseVal := []int{0, 9, 261, 4725, 67509, 831429, 9287109, 97654149, 994388229}
	baseVal := []int{9, 90, 738, 5274, 32490, 168570, 712890, 2345850, 5611770}
	base, cnt := 1, 1
	for n >= base*10 {
		base *= 10
		cnt++
	}
	inverse, mask, t := baseVal[cnt-2]+1, 0, n
	for i := 0; i < cnt; i++ {
		num := t / base
		// 第i位填充比n小的数字, 则后续可以填写任意不重复的数字
		numCnt := 0
		for j := 1; j < num; j++ {
			if mask&(1<<j) == 0 {
				numCnt++
			}
		}
		// 如果不是首位且0未使用，则可以填充0
		if i != 0 && num != 0 && mask&1 == 0 {
			numCnt++
		}
		inverse += numCnt * a(9-i, cnt-i-1)
		if mask&(1<<num) != 0 {
			inverse--
			break
		}
		mask |= 1 << num
		t, base = t%base, base/10
	}
	return n - inverse
}

//leetcode submit region end(Prohibit modification and deletion)

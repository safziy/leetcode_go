package main

import "strconv"

//给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
//
// 如果小数部分为循环小数，则将循环的部分括在括号内。 
//
// 如果存在多个答案，只需返回 任意一个 。 
//
// 对于所有给定的输入，保证 答案字符串的长度小于 10⁴ 。 
//
// 
//
// 示例 1： 
//
// 
//输入：numerator = 1, denominator = 2
//输出："0.5"
// 
//
// 示例 2： 
//
// 
//输入：numerator = 2, denominator = 1
//输出："2"
// 
//
// 示例 3： 
//
// 
//输入：numerator = 4, denominator = 333
//输出："0.(012)"
// 
//
// 
//
// 提示： 
//
// 
// -2³¹ <= numerator, denominator <= 2³¹ - 1 
// denominator != 0 
// 
//
// Related Topics 哈希表 数学 字符串 👍 406 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	var result []byte
	if numerator < 0 != (denominator < 0) {
		result = append(result, '-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	// 整数部分
	integer := abs(numerator) / denominator
	result = append(result, strconv.Itoa(integer)...)
	result = append(result, '.')
	// 小数部分
	circleMap := map[int]int{}
	mod := numerator % denominator
	for mod != 0 && circleMap[mod] == 0 {
		circleMap[mod] = len(result)
		mod *= 10
		result = append(result, '0'+byte(mod/denominator))
		mod = mod % denominator
	}
	// 有循环节
	if mod > 0 {
		insertIndex := circleMap[mod]
		result = append(result[:insertIndex], append([]byte{'('}, result[insertIndex:]...)...)
		result = append(result, ')')
	}
	return string(result)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)

package main

//二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印
//“ERROR”。
//
// 示例1:
//
//
// 输入：0.625
// 输出："0.101"
//
//
// 示例2:
//
//
// 输入：0.1
// 输出："ERROR"
// 提示：0.1无法被二进制准确表示
//
//
//
//
// 提示：
//
//
// 32位包括输出中的 "0." 这两位。
// 题目保证输入用例的小数位数最多只有 6 位
//
//
// 👍 69 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func printBin(num float64) string {
	result := make([]byte, 0, 32)
	result = append(result, '0')
	result = append(result, '.')
	for i := 0; i < 30; i++ {
		num = num * 2
		if num == 0 {
			return string(result)
		} else if num >= 1 {
			result = append(result, '1')
			num -= 1
		} else {
			result = append(result, '0')
		}
	}
	return "ERROR"
}

//leetcode submit region end(Prohibit modification and deletion)

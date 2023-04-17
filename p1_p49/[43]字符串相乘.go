package main
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。 
//
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。 
//
// 
//
// 示例 1: 
//
// 
//输入: num1 = "2", num2 = "3"
//输出: "6" 
//
// 示例 2: 
//
// 
//输入: num1 = "123", num2 = "456"
//输出: "56088" 
//
// 
//
// 提示： 
//
// 
// 1 <= num1.length, num2.length <= 200 
// num1 和 num2 只能由数字组成。 
// num1 和 num2 都不包含任何前导零，除了数字0本身。 
// 
//
// 👍 1181 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			res[i+j+1] += int(num1[i]-'0')*int(num2[j]-'0')
		}
	}
	// 处理进位
	add := 0
	for i := n1+n2-1; i >= 0 ; i-- {
		res[i] += add
		add = res[i] / 10
		res[i] %=  10
	}
	// 消除前置0
	i := 0
	for res[i] == 0 {
		i++
	}
	result := make([]byte, 0, n1+n2-i)
	for ; i <=n1+n2-1 ; i++ {
		result = append(result, byte(res[i] + '0'))
	}
	return string(result)
}
//leetcode submit region end(Prohibit modification and deletion)

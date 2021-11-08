package p50_p99

/*
50. Pow(x, n)
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。


提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/powx-n
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	result := 1.0
	for n > 0 {
		if n%2 != 0 {
			result *= x
		}
		x *= x
		n = n >> 1
	}
	return result
}

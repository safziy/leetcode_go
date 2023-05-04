package main

//你会得到一个字符串 text 。你应该把它分成 k 个子字符串 (subtext1, subtext2，…， subtextk) ，要求满足:
//
//
// subtexti 是 非空 字符串
// 所有子字符串的连接等于 text ( 即subtext1 + subtext2 + ... + subtextk == text )
// 对于所有 i 的有效值( 即 1 <= i <= k ) ，subtexti == subtextk - i + 1 均成立
//
//
// 返回k可能最大值。
//
//
//
// 示例 1：
//
//
//输入：text = "ghiabcdefhelloadamhelloabcdefghi"
//输出：7
//解释：我们可以把字符串拆分成 "(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)"。
//
//
// 示例 2：
//
//
//输入：text = "merchant"
//输出：1
//解释：我们可以把字符串拆分成 "(merchant)"。
//
//
// 示例 3：
//
//
//输入：text = "antaprezatepzapreanta"
//输出：11
//解释：我们可以把字符串拆分成 "(a)(nt)(a)(pre)(za)(tpe)(za)(pre)(a)(nt)(a)"。
//
//
//
//
// 提示：
//
//
// 1 <= text.length <= 1000
// text 仅由小写英文字符组成
//
//
// 👍 108 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestDecomposition(text string) int {
	// 方法一: 动态规划
	//max := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//n := len(text)
	//half := n >> 1
	//// dp[i] 表示以中间字符左右各加i个字符能拆成段式回文的最大k值
	//dp := make([]int, half+1)
	//if n%2 == 1 {
	//	dp[0] = 1
	//}
	//for i := 1; i <= half; i++ {
	//	left, right, res := n>>1-i, (n+1)>>1+i, 1
	//	for j := 1; j <= i; j++ {
	//		if text[left:left+j] == text[right-j:right] {
	//			res = max(res, dp[i-j]+2)
	//		}
	//	}
	//	dp[i] = res
	//}
	//return dp[half]

	// 方法二: 双指针
	n, k := len(text), 0
	left, right := 0, n
	for left < right {
		l, halt := 1, (right-left)>>1
		for l <= halt {
			if text[left:left+l] == text[right-l:right] {
				k += 2
				left, right = left+l, right-l
				break
			}
			l++
		}
		if l > halt {
			k++
			break
		}
	}
	return k
}

//leetcode submit region end(Prohibit modification and deletion)

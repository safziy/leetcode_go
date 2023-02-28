package main

import (
	"math"
)

//电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
//
// 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
//
// 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按
//钮，以此逐个拼写完 key 中的所有字符。
//
// 旋转 ring 拼出 key 字符 key[i] 的阶段中：
//
//
// 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于
//字符 key[i] 。
// 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段
//）, 直至完成所有拼写。
//
//
//
//
// 示例 1：
//
//
//
//
//
//
//
//
//输入: ring = "godding", key = "gd"
//输出: 4
//解释:
// 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
// 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
// 当然, 我们还需要1步进行拼写。
// 因此最终的输出是 4。
//
//
// 示例 2:
//
//
//输入: ring = "godding", key = "godding"
//输出: 13
//
//
//
//
// 提示：
//
//
// 1 <= ring.length, key.length <= 100
// ring 和 key 只包含小写英文字母
// 保证 字符串 key 一定可以由字符串 ring 旋转拼出
//
//
// Related Topics 深度优先搜索 广度优先搜索 字符串 动态规划 👍 248 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findRotateSteps(ring string, key string) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	rLen, kLen := len(ring), len(key)
	// pos[] 代表 每个字符出现在ring的所有位置
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	// dp[i] 代表当前转动到ring上第i个字符花费的最少步数
	dp := make([]int, rLen)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	// 先计算一下转到key的第一个字符需要的步数
	for _, p := range pos[key[0]-'a'] {
		dp[p] = min(p, rLen-p) + 1
	}
	for i := 1; i < kLen; i++ {
		// 当前需要转动到key的第i个字符
		for _, j := range pos[key[i]-'a'] {
			if key[i] == key[i-1] {
				// 如果当前字符和上一个字符相同，则直接步数加1
				dp[j] = dp[j] + 1
				continue
			} else {
				dp[j] = math.MaxInt32
			}
			// 遍历一下ring上相同字符对应的位置
			for _, k := range pos[key[i-1]-'a'] {
				// dp[k] 即为上一个字符可能的位置
				dp[j] = min(dp[j], dp[k]+min(abs(j-k), rLen-abs(j-k))+1)
			}
		}
	}
	// 遍历最后一个字符对应的位置最小值即答案
	result := math.MaxInt32
	for _, k := range pos[key[kLen-1]-'a'] {
		result = min(result, dp[k])
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

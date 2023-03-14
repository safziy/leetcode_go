package main

//给定一个放有字母和数字的数组，找到最长的子数组，且包含的字母和数字的个数相同。
//
// 返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。 
//
// 示例 1: 
//
// 
//输入: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K",
//"L","M"]
//
//输出: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
// 
//
// 示例 2: 
//
// 
//输入: ["A","A"]
//
//输出: []
// 
//
// 提示： 
//
// 
// array.length <= 100000 
// 
//
// 👍 147 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findLongestSubarray(array []string) []string {
	posMap := map[int]int{}
	posMap[0] = -1
	sum := 0
	maxLen, start, end := 0, -1, -1
	for i, s := range array {
		if s[0] >= '0' && s[0] <= '9' {
			sum++
		} else {
			sum--
		}
		if f, ok := posMap[sum]; ok {
			if i-f > maxLen {
				maxLen = i - f
				start, end = f + 1, i
			}
		} else {
			posMap[sum] = i
		}
	}
	if start != -1 {
		return array[start : end+1]
	}
	return []string{}
}

//leetcode submit region end(Prohibit modification and deletion)

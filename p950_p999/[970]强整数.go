package main
//给定三个整数 x 、 y 和 bound ，返回 值小于或等于 bound 的所有 强整数 组成的列表 。 
//
// 如果某一整数可以表示为 xⁱ + yʲ ，其中整数 i >= 0 且 j >= 0，那么我们认为该整数是一个 强整数 。 
//
// 你可以按 任何顺序 返回答案。在你的回答中，每个值 最多 出现一次。 
//
// 
//
// 示例 1： 
//
// 
//输入：x = 2, y = 3, bound = 10
//输出：[2,3,4,5,7,9,10]
//解释： 
//2 = 2⁰ + 3⁰
//3 = 2¹ + 3⁰
//4 = 2⁰ + 3¹
//5 = 2¹ + 3¹
//7 = 2² + 3¹
//9 = 2³ + 3⁰
//10 = 2⁰ + 3² 
//
// 示例 2： 
//
// 
//输入：x = 3, y = 5, bound = 15
//输出：[2,4,6,8,10,14]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= x, y <= 100 
// 0 <= bound <= 10⁶ 
// 
//
// 👍 100 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func powerfulIntegers(x int, y int, bound int) []int {
	res := make(map[int]bool)
	xValue := 1
	for i := 0; i < 21; i++ {
		yValue := 1
		for j := 0; j < 21; j++ {
			value := xValue + yValue
			if value <= bound {
				res[value] = true
			} else {
				break
			}
			yValue *= y
		}
		if xValue > bound {
			break
		}
		xValue *= x
	}
	var result []int
	for k := range res {
		result = append(result, k)
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

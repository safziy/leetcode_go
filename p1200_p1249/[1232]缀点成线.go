package main

//给定一个数组 coordinates ，其中 coordinates[i] = [x, y] ，
// [x, y] 表示横坐标为 x、纵坐标为 y 的点。请你来判断，这些点是否在该坐标系中属于同一条直线上。
//
//
//
// 示例 1：
//
//
//
//
//输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
//输出：true
//
//
// 示例 2：
//
//
//
//
//输入：coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
//输出：false
//
//
//
//
// 提示：
//
//
// 2 <= coordinates.length <= 1000
// coordinates[i].length == 2
// -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
// coordinates 中不含重复的点
//
//
// 👍 129 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	dx, dy := coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]
	for i := 2; i < n; i++ {
		x, y := coordinates[i][0]-coordinates[0][0], coordinates[i][1]-coordinates[0][1]
		if x*dy != y*dx {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

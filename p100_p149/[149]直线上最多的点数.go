package main

import "strconv"

//给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
//
//
//
// 示例 1：
//
//
//输入：points = [[1,1],[2,2],[3,3]]
//输出：3
//
//
// 示例 2：
//
//
//输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= points.length <= 300
// points[i].length == 2
// -10⁴ <= xi, yi <= 10⁴
// points 中的所有点 互不相同
//
//
// 👍 473 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxPoints(points [][]int) int {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	n, result := len(points), 1
	for i := 0; i < n-1; i++ {
		cntMap := map[string]int{}
		for j := i + 1; j < n; j++ {
			x, y := points[i][0]-points[j][0], points[i][1]-points[j][1]
			k := gcd(x, y)
			key := strconv.Itoa(x/k) + "_" + strconv.Itoa(y/k)
			if _, ok := cntMap[key]; ok {
				cntMap[key] = cntMap[key] + 1
			} else {
				cntMap[key] = 2
			}
			if cntMap[key] > result {
				result = cntMap[key]
			}
		}
		if result > n/2 {
			break
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

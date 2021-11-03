package p400_p449

import "container/heap"

/*
407. 接雨水 II
给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。

提示:

m == heightMap.length
n == heightMap[i].length
1 <= m, n <= 200
0 <= heightMap[i][j] <= 2 * 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func trapRainWater(heightMap [][]int) int {
	result := 0
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return result
	}
	vis := make([][]bool, m)
	for r := range vis {
		vis[r] = make([]bool, n)
	}
	minHeap := &MinHeap{}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 || r == m-1 || c == 0 || c == n-1 {
				heap.Push(minHeap, Node{heightMap[r][c], r, c})
				vis[r][c] = true
			}
		}
	}
	dirs := []int{1, 0, -1, 0, 1}
	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(Node)
		for k := 0; k < 4; k++ {
			nr, nc := cur.r+dirs[k], cur.c+dirs[k+1]
			if 0 <= nr && nr < m && 0 <= nc && nc < n && !vis[nr][nc] {
				if heightMap[nr][nc] < cur.h {
					result += cur.h - heightMap[nr][nc]
				}
				heap.Push(minHeap, Node{max(heightMap[nr][nc], cur.h), nr, nc})
				vis[nr][nc] = true
			}
		}
	}
	return result
}

type Node struct {
	h, r, c int
}

type MinHeap []Node

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].h < (*h)[j].h
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(Node))
}

func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

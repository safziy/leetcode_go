package main

import . "leetcode_go/common"

//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
// 
//
// 
// 
// 示例 1： 
// 
// 
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
// 
// 
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 8 
// 
//
// 👍 1463 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	var generateHelper func(start, end int) []*TreeNode
	generateHelper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var allTrees []*TreeNode
		for i := start; i <= end; i++ {
			leftTrees := generateHelper(start, i-1)
			rightTrees := generateHelper(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					root := &TreeNode{Val: i}
					root.Left = left
					root.Right = right
					allTrees = append(allTrees, root)
				}
			}
		}
		return allTrees
	}
	return generateHelper(1, n)
}

//leetcode submit region end(Prohibit modification and deletion)

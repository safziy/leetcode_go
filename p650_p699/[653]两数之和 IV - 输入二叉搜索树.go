package main

import . "leetcode_go/common"
//给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。 
//
// 
//
// 示例 1： 
// 
// 
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
// 
//
// 示例 2： 
// 
// 
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [1, 10⁴]. 
// -10⁴ <= Node.val <= 10⁴ 
// 题目数据保证，输入的 root 是一棵 有效 的二叉搜索树 
// -10⁵ <= k <= 10⁵ 
// 
//
// 👍 465 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	checkMap := map[int]struct{}{}
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool{
		if node == nil {
			return false
		}
		if _, ok := checkMap[k-node.Val]; ok {
			return true
		}
		checkMap[node.Val] = struct{}{}
		if dfs(node.Left) {
			return true
		}
		if dfs(node.Right) {
			return true
		}
		return false
	}
	return dfs(root)
}
//leetcode submit region end(Prohibit modification and deletion)

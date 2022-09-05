package main

import . "leetcode_go/common"
//给定一棵二叉树 root，返回所有重复的子树。 
//
// 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。 
//
// 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：root = [1,2,3,4,null,2,4,null,null,4]
//输出：[[2,4],[4]] 
//
// 示例 2： 
//
// 
//
// 
//输入：root = [2,1,1]
//输出：[[1]] 
//
// 示例 3： 
//
// 
//
// 
//输入：root = [2,2,2,3,null,3,null]
//输出：[[2,3],[3]] 
//
// 
//
// 提示： 
//
// 
// 树中的结点数在[1,10^4]范围内。 
// -200 <= Node.val <= 200 
// 
//
// Related Topics 树 深度优先搜索 哈希表 二叉树 👍 558 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	idx := 0
	var result []*TreeNode
	checkMap := map[[3]int]int{}
	resultMap := map[int]bool{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		check := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
		if checkIdx, ok :=  checkMap[check]; ok {
			if _, ok1 := resultMap[checkIdx]; !ok1 {
				result = append(result, node)
				resultMap[checkIdx] = true
			}
			return checkIdx
		}
		idx ++
		checkMap[check] = idx
		return idx
	}
	dfs(root)
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

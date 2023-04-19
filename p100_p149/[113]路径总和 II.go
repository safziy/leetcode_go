package main

import . "leetcode_go/common"

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//
//
// 示例 2：
//
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1,2], targetSum = 0
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
// 👍 953 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	// 方法一: 深度优先搜索
	var result [][]int
	if root == nil {
		return result
	}
	var path []int
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		targetSum -= cur.Val
		path = append(path, cur.Val)
		if cur.Left == nil && cur.Right == nil {
			if targetSum == 0 {
				result = append(result, append([]int{}, path...))
			}
		}
		if cur.Left != nil {
			dfs(cur.Left)
		}
		if cur.Right != nil {
			dfs(cur.Right)
		}
		targetSum += cur.Val
		path = path[:len(path)-1]
	}
	dfs(root)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

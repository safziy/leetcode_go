package main

import . "leetcode_go/common"
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]
// 
//
// 示例 2: 
//
// 
//输入: [1,null,3]
//输出: [1,3]
// 
//
// 示例 3: 
//
// 
//输入: []
//输出: []
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [0,100] 
// 
// -100 <= Node.val <= 100 
// 
//
// 👍 843 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	// 方法一: 深度优先搜索
	//result := []int{}
	//var dfs func(cur *TreeNode, level int)
	//dfs = func(cur *TreeNode, level int) {
	//	if cur == nil {
	//		return
	//	}
	//	if len(result) <= level {
	//		result = append(result, cur.Val)
	//	} else {
	//		result[level] = cur.Val
	//	}
	//	dfs(cur.Left, level+1)
	//	dfs(cur.Right, level+1)
	//}
	//dfs(root, 0)
	//return result

	// 方法二: 广度优先搜索
	result := []int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		copyQueue := queue
		queue = nil
		for _, node := range copyQueue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

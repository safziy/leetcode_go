package main

import . "leetcode_go/common"

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
// 
//
// 示例 2： 
//
// 
//输入：root = [1]
//输出：[[1]]
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 2000] 内 
// -1000 <= Node.val <= 1000 
// 
//
// 👍 1626 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result, temp := [][]int{}, []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root, nil)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			result = append(result, temp)
			temp = []int{}
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import . "leetcode_go/common"

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
// 👍 2343 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 方法一：递归
	//var isEqual func(tree1, tree2 *TreeNode) bool
	//isEqual = func(tree1, tree2 *TreeNode) bool {
	//	if tree1 == nil && tree2 == nil {
	//		return true
	//	}
	//	if tree1 == nil || tree2 == nil {
	//		return false
	//	}
	//	if tree1.Val != tree2.Val {
	//		return false
	//	}
	//	return isEqual(tree1.Left, tree2.Right) && isEqual(tree1.Right, tree2.Left)
	//}
	//return isEqual(root.Left, root.Right)

	// 方法二：迭代
	queue1, queue2 := []*TreeNode{}, []*TreeNode{}
	queue1 = append(queue1, root.Left)
	queue2 = append(queue2, root.Right)
	for {
		if len(queue1) == 0 && len(queue2) == 0 {
			return true
		}
		if len(queue1) == 0 || len(queue2) == 0 {
			return false
		}
		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1 == nil && node2 == nil {
			continue
		} else if node1 == nil || node2 == nil {
			return false
		} else if node1.Val == node2.Val {
			queue1 = append(queue1, node1.Left, node1.Right)
			queue2 = append(queue2, node2.Right, node2.Left)
		} else {
			return false
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)

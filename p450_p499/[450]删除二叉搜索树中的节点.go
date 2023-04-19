package main

import . "leetcode_go/common"

//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的
//根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。
//
//
//
//
// 示例 1:
//
//
//
//
//输入：root = [5,3,6,2,4,null,7], key = 3
//输出：[5,4,6,2,null,null,7]
//解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
//一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
//另一个正确答案是 [5,2,6,null,4,null,7]。
//
//
//
//
// 示例 2:
//
//
//输入: root = [5,3,6,2,4,null,7], key = 0
//输出: [5,3,6,2,4,null,7]
//解释: 二叉树不包含值为 0 的节点
//
//
// 示例 3:
//
//
//输入: root = [], key = 0
//输出: []
//
//
//
// 提示:
//
//
// 节点数的范围 [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// 节点值唯一
// root 是合法的二叉搜索树
// -10⁵ <= key <= 10⁵
//
//
//
//
// 进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。
//
// 👍 1114 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 方法一: 递归
	if root == nil {
		return nil
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Left == nil || root.Right == nil {
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	} else {
		next := root.Right
		for next.Left != nil {
			next = next.Left
		}
		next.Right = deleteNode(root.Right, next.Val)
		next.Left = root.Left
		return next
	}
	return root

	// 方法二: 迭代
	//var del, pre *TreeNode = root, nil
	//for del != nil && del.Val != key {
	//	if del.Val > key {
	//		pre, del = del, del.Left
	//	} else {
	//		pre, del = del, del.Right
	//	}
	//}
	//// 没有需要删除的节点
	//if del == nil {
	//	return root
	//}
	//if del.Left == nil && del.Right == nil {
	//	del = nil
	//} else if del.Left == nil {
	//	del = del.Right
	//} else if del.Right == nil {
	//	del = del.Left
	//} else {
	//	next, nextParent := del.Right, del
	//	for next.Left != nil {
	//		nextParent = next
	//		next = next.Left
	//	}
	//	if nextParent.Val == del.Val {
	//		nextParent.Right = next.Right
	//	} else {
	//		nextParent.Left = next.Right
	//	}
	//	next.Right = del.Right
	//	next.Left = del.Left
	//	del = next
	//}
	//if pre == nil {
	//	return del
	//}
	//if pre.Left != nil && pre.Left.Val == key {
	//	pre.Left = del
	//} else {
	//	pre.Right = del
	//}
	//return root
}

//leetcode submit region end(Prohibit modification and deletion)

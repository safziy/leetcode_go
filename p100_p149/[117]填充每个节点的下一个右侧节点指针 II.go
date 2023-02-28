package main

import . "leetcode_go/common"

//给定一个二叉树：
//
//
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
//
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
//
// 初始状态下，所有 next 指针都被设置为 NULL 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,4,5,null,7]
//输出：[1,#,2,3,#,4,5,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指
//针连接），'#' 表示每层的末尾。
//
// 示例 2：
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
// 树中的节点数在范围 [0, 6000] 内
// -100 <= Node.val <= 100
//
//
// 进阶：
//
//
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。
//
//
//
//
//
// 👍 691 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect2(root *Node) *Node {
	cur, nextFirst := root, (*Node)(nil)
	for cur != nil {
		if nextFirst == nil {
			if cur.Left != nil {
				nextFirst = cur.Left
			} else if cur.Right != nil {
				nextFirst = cur.Right
			}
		}
		childNext := (*Node)(nil)
		next := cur.Next
		for next != nil {
			if next.Left != nil {
				childNext = next.Left
				break
			} else if next.Right != nil {
				childNext = next.Right
				break
			}
			next = next.Next
		}
		if cur.Left != nil {
			if cur.Right != nil {
				cur.Left.Next = cur.Right
			} else if cur.Next != nil {
				cur.Left.Next = childNext
			}
		}
		if cur.Right != nil {
			cur.Right.Next = childNext
		}
		if cur.Next != nil {
			cur = cur.Next
		} else {
			cur = nextFirst
			nextFirst = nil
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

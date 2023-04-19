package main

import . "leetcode_go/common"

//给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
//
//
// （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
//输出：7
//解释：
//我们有大量的节点与其祖先的差值，其中一些如下：
//|8 - 3| = 5
//|3 - 7| = 4
//|8 - 1| = 7
//|10 - 13| = 3
//在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。
//
//
// 示例 2：
//
//
//输入：root = [1,null,2,null,0,3]
//输出：3
//
//
//
//
// 提示：
//
//
// 树中的节点数在 2 到 5000 之间。
// 0 <= Node.val <= 10⁵
//
//
// 👍 162 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	result := 0
	var dfs func(cur *TreeNode, maxVal, minVal int)
	dfs = func(cur *TreeNode, maxVal, minVal int) {
		if cur == nil {
			return
		}
		diff := max(abs(maxVal-cur.Val), abs(minVal-cur.Val))
		result = max(diff, result)
		newMaxVal := max(cur.Val, maxVal)
		newMinVal := min(cur.Val, minVal)
		dfs(cur.Left, newMaxVal, newMinVal)
		dfs(cur.Right, newMaxVal, newMinVal)
	}
	dfs(root, root.Val, root.Val)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

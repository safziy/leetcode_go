package main

import (
	. "leetcode_go/common"
)

//
//
// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则
//，返回 false 。
//
//
//
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
//
//
//
// 示例 1：
//
//
//输入：root = [3,4,5,1,2], subRoot = [4,1,2]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// root 树上的节点数量范围是 [1, 2000]
// subRoot 树上的节点数量范围是 [1, 1000]
// -10⁴ <= root.val <= 10⁴
// -10⁴ <= subRoot.val <= 10⁴
//
//
// 👍 874 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	const nilVal = 1024
	var dfs func(n *TreeNode) []int
	dfs = func(n *TreeNode) []int {
		if n == nil {
			return []int{nilVal}
		}
		res := []int{}
		res = append(res, n.Val)
		res = append(res, dfs(n.Left)...)
		res = append(res, dfs(n.Right)...)
		return res
	}
	src := dfs(root)
	target := dfs(subRoot)
	return kmp(src, target)
}

func kmp(src, target []int) bool {
	n, m := len(src), len(target)
	dp := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && target[i] != target[j] {
			j = dp[j-1]
		}
		if target[i] == target[j] {
			j++
		}
		dp[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && src[i] != target[j] {
			j = dp[j-1]
		}
		if src[i] == target[j] {
			j++
		}
		if j == m {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

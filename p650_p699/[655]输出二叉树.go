package main

import (
	"strconv"
)

//给你一棵二叉树的根节点 root ，请你构造一个下标从 0 开始、大小为 m x n 的字符串矩阵 res ，用以表示树的 格式化布局 。构造此格式化布局矩
//阵需要遵循以下规则： 
//
// 
// 树的 高度 为 height ，矩阵的行数 m 应该等于 height + 1 。 
// 矩阵的列数 n 应该等于 2ʰᵉⁱᵍʰᵗ⁺¹ - 1 。 
// 根节点 需要放置在 顶行 的 正中间 ，对应位置为 res[0][(n-1)/2] 。 
// 对于放置在矩阵中的每个节点，设对应位置为 res[r][c] ，将其左子节点放置在 res[r+1][c-2ʰᵉⁱᵍʰᵗ⁻ʳ⁻¹] ，右子节点放置在 
//res[r+1][c+2ʰᵉⁱᵍʰᵗ⁻ʳ⁻¹] 。 
// 继续这一过程，直到树中的所有节点都妥善放置。 
// 任意空单元格都应该包含空字符串 "" 。 
// 
//
// 返回构造得到的矩阵 res 。 
//
// 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [1,2]
//输出：
//[["","1",""],
// ["2","",""]]
// 
//
// 示例 2： 
// 
// 
//输入：root = [1,2,3,null,4]
//输出：
//[["","","","1","","",""],
// ["","2","","","","3",""],
// ["","","4","","","",""]]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数在范围 [1, 2¹⁰] 内 
// -99 <= Node.val <= 99 
// 树的深度在范围 [1, 10] 内 
// 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 178 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	height := 0
	dfsHeight(root, 1, func(r int) {
		if r > height {
			height = r
		}
	})
	result := make([][]string, height)
	for i := 0; i < height; i++ {
		result[i] = make([]string, (1<<height)-1)
	}
	dfsPrint(root, height, 1, 1 << (height-1), func(r, c, v int) {
		result[r-1][c-1] = strconv.Itoa(v)
	})
	return result
}

func dfsHeight(node *TreeNode, r int, fuc func(r int)) {
	fuc(r)
	if node.Left != nil {
		dfsHeight(node.Left, r+1,  fuc)
	}
	if node.Right != nil {
		dfsHeight(node.Right, r+1,  fuc)
	}
}

func dfsPrint(node *TreeNode, h, r, c int, fuc func(r, c, v int)) {
	fuc(r, c, node.Val)
	if node.Left != nil {
		dfsPrint(node.Left, h, r+1, c - (1 << (h-r-1)), fuc)
	}
	if node.Right != nil {
		dfsPrint(node.Right, h, r+1, c + (1 << (h-r-1)), fuc)
	}
}
//leetcode submit region end(Prohibit modification and deletion)

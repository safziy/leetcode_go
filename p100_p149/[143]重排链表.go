package main

import . "leetcode_go/common"

//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
//L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [1,2,3,4]
//输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5]
//输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
//
// 👍 1186 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMidNode(head)
	tail := reverse(mid)
	merge(head, tail)
}

func findMidNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		head = slow
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			head.Next = nil
			return slow
		}
		fast = fast.Next
	}
	head.Next = nil
	return slow
}

func reverse(head *ListNode) *ListNode {
	cur, next, pre := head, (*ListNode)(nil), (*ListNode)(nil)
	for cur != nil {
		next, cur.Next = cur.Next, pre
		pre, cur = cur, next
	}
	return pre
}

func merge(first, second *ListNode) {
	firstNext, secondNext := (*ListNode)(nil), (*ListNode)(nil)
	for first != nil && second != nil {
		firstNext = first.Next
		secondNext = second.Next
		first.Next = second
		second.Next = firstNext
		first = firstNext
		second = secondNext
	}
}

//leetcode submit region end(Prohibit modification and deletion)

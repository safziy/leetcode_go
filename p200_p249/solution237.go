package p200_p249

import . "leetcode_go/common/singly_linked"

/*
237. 删除链表中的节点

请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点head ，只能直接访问 要被删除的节点 。

题目数据保证需要删除的节点 不是末尾节点 。

提示：

链表中节点的数目范围是 [2, 1000]
-1000 <= Node.val <= 1000
链表中每个节点的值都是唯一的
需要删除的节点 node 是 链表中的一个有效节点 ，且 不是末尾节点

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	if next.Next != nil {
		deleteNode(next)
	} else {
		node.Next = nil
	}
}

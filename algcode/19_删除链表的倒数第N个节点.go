package algcode

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？


*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	p := head
	for p != nil {
		count++
		p = p.Next
	}
	if n > count {
		return head
	}
	if count == n {
		head = head.Next
		return head
	}
	cnt := count - n - 1
	i := 0
	p = head
	for i < cnt {
		p = p.Next
		i++
	}
	p.Next = p.Next.Next
	return head
}

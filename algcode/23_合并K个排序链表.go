package algcode

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6


*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		lists = MergeLong2Short(lists)
	}

	return lists[0]
}

func MergeLong2Short(lists []*ListNode) (newList []*ListNode) {

	l := 0
	r := len(lists) - 1

	for l < r {
		lists[l] = merge(lists[l], lists[r])
		l++
		r--
	}
	if l == r {
		l++
	}
	return lists[:l]
}

func merge(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	Head := &ListNode{}
	p := Head

	for a != nil && b != nil {
		if a.Val < b.Val {
			p.Next = b
			p = b
			b = b.Next
		} else {
			p.Next = a
			p = a
			a = a.Next
		}

		if a == nil {
			p.Next = b
		} else if b == nil {
			p.Next = a
		}

	}
	return Head.Next
}

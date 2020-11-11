package leetcode

import (
	"leetcode-everyday/structures"
)

// ListNode ...
type ListNode = structures.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	l := res //注意这一条语句 他是用他来做循环 然后最后返回res.Next
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 == nil {
		l.Next = l2
	}
	if l2 == nil {
		l.Next = l1
	}
	return res.Next
}

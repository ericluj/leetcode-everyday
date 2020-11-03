package leetcode

import (
	"leetcode-everyday/structures"
)

// ListNode ...
type ListNode = structures.ListNode

// ReverseLinkedList ...
func ReverseLinkedList(pHead *ListNode) *ListNode {
	var pre *ListNode
	cur := pHead
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

package leetcode

import (
	s "leetcode-everyday/structures"
)

// HasCycle ...
func HasCycle(head *s.ListNode) bool {
	first := head
	second := head
	for first != nil && first.Next != nil { //快指针做判断是否越界
		first = first.Next.Next
		second = second.Next
		if first == second {
			return true
		}
	}
	return false
}

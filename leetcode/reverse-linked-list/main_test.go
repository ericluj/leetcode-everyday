package leetcode

import (
	"testing"

	s "leetcode-everyday/structures"
)

func Test_ReverseLinkedList(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	out := []int{5, 4, 3, 2, 1}
	res := s.List2Ints(ReverseLinkedList(s.Ints2List(in)))
	if !s.IntSliceEqual(out, res) {
		t.Errorf("错误")
	}
	return
}

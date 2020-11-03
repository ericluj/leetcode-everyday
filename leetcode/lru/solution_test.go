package leetcode

import (
	"testing"

	s "leetcode-everyday/structures"
)

func Test_LRU(t *testing.T) {
	in := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	out := []int{1, -1}
	res := LRU(in, 3)
	if !s.IntSliceEqual(out, res) {
		t.Errorf("错误")
	}
	return
}

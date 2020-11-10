package leetcode

import (
	s "leetcode-everyday/structures"
)

func threeOrders(root *s.TreeNode) [][]int {
	var (
		pre  []int
		mid  []int
		last []int
	)
	PreOrder(root, &pre)
	MidOrder(root, &mid)
	LastOrder(root, &last)
	return [][]int{pre, mid, last}
}

// PreOrder ...
func PreOrder(root *s.TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	PreOrder(root.Left, nums)
	PreOrder(root.Right, nums)
}

// MidOrder ...
func MidOrder(root *s.TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	MidOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	MidOrder(root.Right, nums)
}

// LastOrder ...
func LastOrder(root *s.TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	LastOrder(root.Left, nums)
	LastOrder(root.Right, nums)
	*nums = append(*nums, root.Val)
}

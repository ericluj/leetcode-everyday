package leetcode

// UpperBound ...
func UpperBound(n int, v int, a []int) int {
	if n == 0 { //防止下面的判断越界
		return 1
	}
	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] < v {
			left = mid + 1
		} else if a[mid] > v {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}
	if left >= n || a[left] != v { //判断不存在数字
		if a[0] > v { //注意审题，第一个大于等于数，那么没有找到且比他小返回第一个数就好了
			return 1
		}
		return n + 1
	}
	return left + 1 //输出未知从1开始计算所以加1
}

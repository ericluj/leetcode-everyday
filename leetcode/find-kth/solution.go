package leetcode

func findKth(a []int, n int, K int) int {
	quickSort(a, 0, n-1)
	return a[n-K]
}

func quickSort(a []int, st, ed int) {
	if st >= ed {
		return
	}
	x := a[st] //基准数
	i := st
	j := ed
	for i != j {
		for a[j] >= x && i < j {
			j--
		}
		for a[i] <= x && i < j {
			i++
		}
		a[i], a[j] = a[j], a[i]
	}
	a[st], a[i] = a[i], a[st]
	quickSort(a, st, i-1)
	quickSort(a, i+1, ed)
}

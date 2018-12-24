package t006

// QuickSort 快速排序
func QuickSort(a []int, n int) {
	quickSort(a, 0, n-1)
}

func quickSort(a []int, p, r int) {
	if p >= r {
		return
	}

	// 找到分区点
	q := partition(a, p, r)
	quickSort(a, p, q-1)
	quickSort(a, q+1, r)
}

// 分区函数
func partition(a []int, p, r int) int {
	// 一般情况使用分区区间的最后一个元素作为分区点
	pivot := a[r]
	i := p

	for j := p; j < r; j++ {
		if a[j] < pivot { // 把小于分区点的元素放到左侧
			tmp := a[j]
			a[j] = a[i]
			a[i] = tmp
			i = i + 1
		}
	}

	if i != r {
		tmp := a[i]
		a[i] = a[r]
		a[r] = tmp
	}

	return i
}

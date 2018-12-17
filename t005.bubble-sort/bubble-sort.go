package t005

// BubbleSort 冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}

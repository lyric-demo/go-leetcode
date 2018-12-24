package t006

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{3, 2, 1, 6, 5, 4, 9, 8, 7}
	QuickSort(a, 9)

	target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(a, target) {
		t.Error("无效的结果值：", a)
	}
}

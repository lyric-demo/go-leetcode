package t005

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{6, 5, 1, 2, 3, 4}
	BubbleSort(a)

	target := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(a, target) {
		t.Error("不是期望的结果：", a)
	}
}

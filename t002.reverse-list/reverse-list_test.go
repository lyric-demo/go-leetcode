package reverselist

import "testing"

func TestReverseList(t *testing.T) {
	l1 := ToListNode(1, 2, 3, 4, 5, 6)
	l2 := ReverseList(l1)

	tl := ToListNode(6, 5, 4, 3, 2, 1)
	if !EqualList(l2, tl) {
		t.Error("不是期望的结果：", l2)
		return
	}
}

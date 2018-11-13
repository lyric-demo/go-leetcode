package addtwonumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	// Output: 7 -> 0 -> 8
	l1 := ToListNode(2, 4, 3)
	l2 := ToListNode(5, 6, 4)

	r := AddTwoNumbers(l1, l2)
	if !EqualListNode(r, ToListNode(7, 0, 8)) {
		t.Error("错误的执行结果：", r)
	}
}

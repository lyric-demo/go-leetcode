package addtwonumbers

// AddTwoNumbers 增加两个链表
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	r := NewListNode(0)
	t := r
	sum := 0
	c := 0

	for l1 != nil || l2 != nil || c > 0 {
		sum = c

		if l1 != nil {
			sum += l1.Value
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Value
			l2 = l2.Next
		}
		c = sum / 10

		t.Next = NewListNode(sum % 10)
		t = t.Next
	}

	return r.Next
}

package reverselist

// ReverseList 反转链表
func ReverseList(list *ListNode) *ListNode {
	var pre *ListNode
	curr := list

	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre
}

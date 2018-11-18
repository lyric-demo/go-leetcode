package palindromelist

// CheckPalindromeList 检查链表是否是回文串
func CheckPalindromeList(l *ListNode) bool {
	if l == nil {
		return false
	}

	if l.Next == nil {
		return true
	}

	var pre *ListNode
	var curr = l
	var qcurr = l

	for qcurr != nil && qcurr.Next != nil {
		qcurr = qcurr.Next.Next

		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	if qcurr != nil {
		curr = curr.Next
	}

	for curr != nil {
		if pre.Value != curr.Value {
			return false
		}
		curr = curr.Next
		pre = pre.Next
	}

	return true
}

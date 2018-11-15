package reverselist

// ListNode 链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

// NewListNode 创建链表节点
func NewListNode(v int) *ListNode {
	return &ListNode{Value: v}
}

// ToListNode 转换为链表节点
func ToListNode(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	root := NewListNode(vals[0])
	tmp := root

	for i := 1; i < len(vals); i++ {
		tmp.Next = NewListNode(vals[i])
		tmp = tmp.Next
	}

	return root
}

// EqualList 比较链表
func EqualList(l1, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil {
			return false
		}

		if l1.Value != l2.Value {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}

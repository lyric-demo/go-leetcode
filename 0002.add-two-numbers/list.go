package addtwonumbers

// ListNode 链表
type ListNode struct {
	Value int
	Next  *ListNode
}

// NewListNode 创建链表节点
func NewListNode(v int) *ListNode {
	return &ListNode{
		Value: v,
	}
}

// ToListNode 转换为链表
func ToListNode(vals ...int) *ListNode {
	root := NewListNode(0)
	temp := root

	for _, v := range vals {
		temp.Next = NewListNode(v)
		temp = temp.Next
	}

	if root.Next != nil {
		root = root.Next
	}

	return root
}

// EqualListNode 对比链表
func EqualListNode(l1, l2 *ListNode) bool {
	for l1 != nil {
		if l2 == nil || l1.Value != l2.Value {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

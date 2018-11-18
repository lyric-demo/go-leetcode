package palindromelist

// ListNode 单向链表节点
type ListNode struct {
	Value byte
	Next  *ListNode
}

// ToListNode 转换为单向链表
func ToListNode(s ...byte) *ListNode {
	if len(s) == 0 {
		return nil
	}

	root := &ListNode{Value: s[0]}
	tmp := root

	for i := 1; i < len(s); i++ {
		tmp.Next = &ListNode{Value: s[i]}
		tmp = tmp.Next
	}

	return root
}

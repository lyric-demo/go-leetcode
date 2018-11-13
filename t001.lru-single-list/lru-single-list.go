package lrusinglelist

/* 使用单项链表实现LRU缓存淘汰算法
实现思路：
我们维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。
当有一个新的数据被访问时，我们从链表头开始顺序遍历链表。
*/

// ListNode 链表节点
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// NewListNode 创建一个链表节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{Value: v}
}

// SingleList 单向链表
type SingleList struct {
	root   *ListNode
	length int
}

// NewSingleList 创建一个单向链表
func NewSingleList() *SingleList {
	return &SingleList{
		root: NewListNode(0),
	}
}

// Front 获取最前面的节点
func (s *SingleList) Front() *ListNode {
	return s.root.Next
}

// Back 获取最后面的节点
func (s *SingleList) Back() *ListNode {
	cur := s.root.Next

	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}

	return cur
}

// MoveToFront 将节点移动到最前面
func (s *SingleList) MoveToFront(e *ListNode) {
	f := s.Front()
	if f == e {
		return
	}

	cur := f
	for cur != nil {
		next := cur.Next
		if next == e {
			cur.Next = e.Next
			e.Next = f
			s.root.Next = e
			break
		}
		cur = next
	}

}

// InsertBefore 插入到节点之前
func (s *SingleList) InsertBefore(v interface{}, e *ListNode) *ListNode {
	f := s.Front()
	if f == e {
		n := NewListNode(v)
		n.Next = f
		s.root.Next = n
		s.length = s.length + 1
		return n
	}

	cur := f
	for cur != nil {
		next := cur.Next
		if next == e {
			n := NewListNode(v)
			n.Next = e
			cur.Next = n
			s.length = s.length + 1
			return n
		}
		cur = next
	}

	return nil
}

// Remove 删除节点
func (s *SingleList) Remove(e *ListNode) interface{} {
	cur := s.Front()

	for cur != nil {
		if cur == e {
			s.length = s.length - 1
			cur.Next = e.Next
			return cur.Value
		}
		cur = cur.Next
	}

	return nil
}

// Len 获取链表的长度
func (s *SingleList) Len() int {
	return s.length
}

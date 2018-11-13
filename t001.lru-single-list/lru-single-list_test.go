package lrusinglelist

import "testing"

func TestSingleList(t *testing.T) {
	list := NewSingleList()

	node := list.InsertBefore("foo", list.Front())
	if node == nil || node.Value != "foo" || list.Len() != 1 {
		t.Error("节点错误：", node)
		return
	}

	node2 := list.InsertBefore("bar", node)
	if node2 == nil || node2.Value != "bar" || list.Len() != 2 {
		t.Error("节点错误：", node2)
		return
	}

	node3 := list.InsertBefore("la", node)
	if node3 == nil || node3.Value != "la" || list.Len() != 3 {
		t.Error("节点错误：", node3)
		return
	}

	fnode := list.Front()
	if fnode == nil || fnode.Value != "bar" {
		t.Error("节点错误：", fnode)
		return
	}

	bnode := list.Back()
	if bnode == nil || bnode.Value != "foo" {
		t.Error("节点错误：", bnode)
		return
	}

	list.MoveToFront(bnode)

	fnode2 := list.Front()
	if fnode2 == nil || fnode2.Value != "foo" {
		t.Error("节点错误：", fnode2)
		return
	}

	bnode2 := list.Back()
	if bnode2 == nil || bnode2.Value != "la" {
		t.Error("节点错误：", bnode2)
		return
	}

	bv := list.Remove(bnode2)
	if bv != "la" || list.Len() != 2 {
		t.Error("节点错误：", bv)
		return
	}

}

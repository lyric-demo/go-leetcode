package arraystack

import "testing"

func TestArrayStack(t *testing.T) {
	stack := NewArrayStack(3)
	b := stack.Push("1")
	if !b {
		t.Error("无效的返回值：", b)
		return
	}

	b = stack.Push("2")
	if !b {
		t.Error("无效的返回值：", b)
		return
	}

	b = stack.Push("3")
	if !b {
		t.Error("无效的返回值：", b)
		return
	}

	s := stack.Pop()
	if s != "3" {
		t.Error("无效的返回值：", s)
	}

	s = stack.Pop()
	if s != "2" {
		t.Error("无效的返回值：", s)
	}

	s = stack.Pop()
	if s != "1" {
		t.Error("无效的返回值：", s)
	}
}

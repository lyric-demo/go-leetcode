package palindromelist

import "testing"

func TestCheckPalindromeList(t *testing.T) {
	s := []byte("hh")
	if v := CheckPalindromeList(ToListNode(s...)); !v {
		t.Errorf("不是期望的结果：%s - %v", s, v)
		return
	}

	s = []byte("h1h")
	if v := CheckPalindromeList(ToListNode(s...)); !v {
		t.Errorf("不是期望的结果：%s - %v", s, v)
		return
	}

	s = []byte("h11h")
	if v := CheckPalindromeList(ToListNode(s...)); !v {
		t.Errorf("不是期望的结果：%s - %v", s, v)
		return
	}

	s = []byte("heleh")
	if v := CheckPalindromeList(ToListNode(s...)); !v {
		t.Errorf("不是期望的结果：%s - %v", s, v)
		return
	}

	s = []byte("helleh")
	if v := CheckPalindromeList(ToListNode(s...)); !v {
		t.Errorf("不是期望的结果：%s - %v", s, v)
		return
	}
}

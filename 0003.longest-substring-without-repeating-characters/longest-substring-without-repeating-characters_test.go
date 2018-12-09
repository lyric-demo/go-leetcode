package leetcode0003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	l := LengthOfLongestSubstring("abcabcbb")
	if l != 3 {
		t.Error("不是期望的值：", l)
	}

	l = LengthOfLongestSubstring("bbbbbb")
	if l != 1 {
		t.Error("不是期望的值：", l)
	}

	l = LengthOfLongestSubstring("abcdefffffffg")
	if l != 6 {
		t.Error("不是期望的值：", l)
	}
}

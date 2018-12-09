package leetcode0003

// LengthOfLongestSubstring 给定一个字符串，找到最长子字符串的长度而不重复字符
func LengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)

	left, maxLen := 0, 0
	for i, r := range s {
		if v, ok := m[r]; ok && v >= left {
			left = v + 1
		} else if m := i + 1 - left; m > maxLen {
			maxLen = m
		}

		m[r] = i
	}

	return maxLen
}

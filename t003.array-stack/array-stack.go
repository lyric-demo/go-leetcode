package arraystack

// ArrayStack 顺序栈
type ArrayStack struct {
	items []string
	count int
	n     int
}

// NewArrayStack 创建一个顺序栈
func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		items: make([]string, n),
		n:     n,
	}
}

// Push 入栈操作（如果栈满，返回false）
func (s *ArrayStack) Push(item string) bool {
	if s.count == s.n {
		return false
	}

	s.items[s.count] = item
	s.count++

	return true
}

// Pop 出栈操作（如果栈中没有元素，则返回空）
func (s *ArrayStack) Pop() string {
	if s.count == 0 {
		return ""
	}

	item := s.items[s.count-1]
	s.count--

	return item
}

package stack

type node struct {
	data interface{}
	next *node
}

type chainStack struct {
	length int
	top    *node
}

func newChainStack() *chainStack {
	return &chainStack{}
}

func (s *chainStack) Reset() {
	*s = chainStack{
		length: 0,
		top:    nil,
	}
}

func (s *chainStack) Len() int {
	return s.length
}

func (s *chainStack) Top() (interface{}, bool) {
	if s.length == 0 {
		return nil, false
	}
	return s.top.data, true
}

func (s *chainStack) Push(ele interface{}) {
	s.top = &node{
		data: ele,
		next: s.top,
	}
	s.length += 1
	return
}

func (s *chainStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	ele := s.top.data
	s.top.data = nil
	s.top = s.top.next
	s.length -= 1
	return ele
}

func (s *chainStack) Each(handler func(ele interface{}) bool) {
	if handler == nil || s.length == 0 {
		return
	}
	p := s.top
	for p != nil {
		if handler(p.data) {
			return
		}
		p = p.next
	}
}

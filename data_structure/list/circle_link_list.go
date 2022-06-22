package list

import "github.com/pyihe/go-pkg/maths"

type circleNode struct {
	data interface{} // 数据
	pre  *circleNode // 前驱节点
	next *circleNode // 后继节点
}

type circleList struct {
	length int         // 链表长度
	guard  *circleNode // 哨兵节点
}

func newCircleList() *circleList {
	guard := &circleNode{}
	guard.pre = guard
	guard.next = guard
	return &circleList{
		length: 0,
		guard:  guard,
	}
}

func (s *circleList) Len() int {
	return s.length
}

func (s *circleList) Reset() {
	s.guard.data = nil
	s.guard.pre = s.guard
	s.guard.next = s.guard
	*s = circleList{
		length: 0,
	}
}

func (s *circleList) Get(pos int) (interface{}, bool) {
	n := maths.Abs(pos)
	if s.length == 0 || pos == 0 || n > s.length {
		return nil, false
	}

	p := s.guard
	for i := 1; i <= n; i++ {
		switch {
		case pos > 0:
			p = p.next
		default:
			p = p.pre
		}
		if p == s.guard {
			return nil, false
		}
	}
	return p.data, true
}

func (s *circleList) Pos(ele interface{}) int {
	if s.length == 0 {
		return 0
	}
	p := s.guard.next
	for i := 1; i <= s.length; i++ {
		if p.data == ele {
			return i
		}
	}
	return 0
}

func (s *circleList) Insert(pos int, ele interface{}) error {
	n, p := maths.Abs(pos), s.guard.next
	if n == 0 || n > s.length+1 {
		return nil
	}
	for i := 1; i < n; i++ {
		switch {
		case pos > 0:
			p = p.next
		default:
			p = p.pre
		}
	}
	node := &circleNode{
		data: ele,
	}
	switch {
	case pos > 0:
		node.pre = p
		node.next = p.next
		p.next = node

	default:
		node.next = p
		node.pre = p.pre
		p.pre = node
	}
	s.length += 1
	return nil
}

func (s *circleList) Delete(pos int) bool {
	n := maths.Abs(pos)
	if s.length == 0 || n > s.length {
		return false
	}
	p := s.guard
	for i := 1; i < n; i++ {
		switch {
		case pos > 0:
			p = p.next
		default:
			p = p.pre
		}
	}
	switch {
	case pos > 0:
		p.next.data = nil
		p.next = p.next.next
		p.next.next.pre = p
	default:
		p.pre.data = nil
		p.pre = p.pre.pre
		p.pre.pre.next = p
	}
	s.length -= 1
	return true
}

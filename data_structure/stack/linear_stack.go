package stack

const (
	defaultStackLen = 256
)

type linearStack struct {
	length int
	data   []interface{}
}

func newLinearStack(size int) *linearStack {
	if size <= 0 {
		size = defaultStackLen
	}
	s := &linearStack{
		length: 0,
		data:   make([]interface{}, 0, size),
	}
	return s
}

func (s *linearStack) Reset() {
	*s = linearStack{
		data:   make([]interface{}, 0, defaultStackLen),
		length: 0,
	}
}

func (s *linearStack) Len() int {
	return s.length
}

func (s *linearStack) Top() (interface{}, bool) {
	if s.length > 0 {
		return s.data[s.length-1], true
	}
	return nil, false
}

func (s *linearStack) Push(ele interface{}) {
	s.data = append(s.data, ele)
	s.length += 1
}

func (s *linearStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	ele := s.data[s.length-1]
	s.data[s.length-1] = nil
	s.data = s.data[:s.length-1]
	s.length -= 1
	return ele
}

func (s *linearStack) Each(handler func(ele interface{}) bool) {
	if handler == nil || s.length == 0 {
		return
	}
	for i := s.length - 1; i >= 0; i-- {
		if handler(s.data[i]) {
			return
		}
	}
}

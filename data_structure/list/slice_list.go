package list

const (
	size = 256
)

// 切片线性表
type sliceList struct {
	data   []interface{}
	length int
}

func newSliceList() Lister {
	sl := &sliceList{
		data:   make([]interface{}, 0, size),
		length: 0,
	}
	return sl
}

func (s *sliceList) Len() int {
	return s.length
}

func (s *sliceList) Reset() {
	*s = sliceList{
		data:   make([]interface{}, 0, size),
		length: 0,
	}
}

func (s *sliceList) Get(pos int) (interface{}, bool) {
	// 如果pos不处于线性表元素范围内
	if s.length == 0 || pos < 1 || pos > s.length {
		return nil, false
	}
	return s.data[pos-1], true
}

func (s *sliceList) Pos(ele interface{}) int {
	for i, e := range s.data {
		if e == ele {
			return i + 1
		}
	}
	return 0
}

func (s *sliceList) Insert(pos int, data interface{}) error {
	pos -= 1
	if pos < 0 || pos > s.length {
		return ErrInvalidPos
	}

	s.data = append(s.data, nil)
	switch {
	case pos < s.length:
		copy(s.data[pos+1:], s.data[pos:s.length+1])
		s.data[pos] = data
	default:
		s.data[s.length] = data
	}
	s.length += 1
	return nil
}

func (s *sliceList) Delete(pos int) bool {
	pos -= 1
	if pos < 0 || pos >= s.length {
		return false
	}
	copy(s.data[pos:], s.data[pos+1:])
	s.data[s.length-1] = nil
	s.data = s.data[:s.length-1]
	s.length -= 1
	return true
}

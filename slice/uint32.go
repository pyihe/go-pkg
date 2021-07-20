package slice

import "fmt"

type uint32Slice []uint32

func newUint32Slice(c int) *uint32Slice {
	s := make(uint32Slice, 0, c)
	return &s
}

func (i32 *uint32Slice) String() string {
	if i32 == nil {
		return ""
	}
	return fmt.Sprintf("%v", *i32)
}

func (i32 *uint32Slice) PushBack(x interface{}) (bool, int) {
	if i32 == nil {
		return false, 0
	}
	v, ok := x.(uint32)
	if !ok {
		return false, 0
	}
	n := len(*i32)
	c := cap(*i32)
	if n+1 > c {
		ns := make(uint32Slice, n, c*2)
		copy(ns, *i32)
		*i32 = ns
	}
	*i32 = (*i32)[0 : n+1]
	(*i32)[n] = v
	return true, n + 1
}

func (i32 *uint32Slice) PushFront(x interface{}) (bool, int) {
	if i32 == nil {
		return false, 0
	}
	v, ok := x.(uint32)
	if !ok {
		return false, 0
	}
	n := len(*i32)
	c := cap(*i32)
	if n+1 > c {
		ns := make(uint32Slice, n, c*2)
		copy(ns, *i32)
		*i32 = ns
	}
	*i32 = (*i32)[0 : n+1]
	copy((*i32)[1:n+1], (*i32)[0:n])
	(*i32)[0] = v
	return true, n + 1
}

func (i32 *uint32Slice) PopBack() (bool, interface{}) {
	if i32 == nil {
		return false, nil
	}
	n := len(*i32)
	c := cap(*i32)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uint32Slice, n, c/2)
		copy(ns, *i32)
		*i32 = ns
	}
	x := (*i32)[n-1]
	*i32 = (*i32)[0 : n-1]
	return true, x
}

func (i32 *uint32Slice) PopFront() (bool, interface{}) {
	if i32 == nil {
		return false, nil
	}
	n := len(*i32)
	c := cap(*i32)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uint32Slice, n, c/2)
		copy(ns, *i32)
		*i32 = ns
	}
	x := (*i32)[0]
	*i32 = (*i32)[1:n]
	return true, x
}

func (i32 *uint32Slice) Index(x interface{}) (index int) {
	index = -1
	if i32 == nil {
		return
	}
	v, ok := x.(uint32)
	if !ok {
		return
	}

	for i := range *i32 {
		if (*i32)[i] == v {
			index = i
			break
		}
	}
	return
}

func (i32 *uint32Slice) IndexValue(index int) interface{} {
	if i32 == nil {
		return nil
	}
	n := len(*i32)
	if index < 0 || index >= n {
		return nil
	}
	return (*i32)[index]
}

func (i32 *uint32Slice) Range(fn func(index int, ele interface{}) bool) {
	if i32 == nil {
		return
	}
	for i, x := range *i32 {
		if fn(i, x) {
			break
		}
	}
}

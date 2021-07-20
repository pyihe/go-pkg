package slice

import "fmt"

type uint64Slice []uint64

func newUint64Slice(c int) *uint64Slice {
	s := make(uint64Slice, 0, c)
	return &s
}

func (i64 *uint64Slice) String() string {
	if i64 == nil {
		return ""
	}
	return fmt.Sprintf("%v", *i64)
}

func (i64 *uint64Slice) PushBack(x interface{}) (bool, int) {
	if i64 == nil {
		return false, 0
	}
	v, ok := x.(uint64)
	if !ok {
		return false, 0
	}
	n := len(*i64)
	c := cap(*i64)
	if n+1 > c {
		ns := make(uint64Slice, n, c*2)
		copy(ns, *i64)
		*i64 = ns
	}
	*i64 = (*i64)[0 : n+1]
	(*i64)[n] = v
	return true, n + 1
}

func (i64 *uint64Slice) PushFront(x interface{}) (bool, int) {
	if i64 == nil {
		return false, 0
	}
	v, ok := x.(uint64)
	if !ok {
		return false, 0
	}
	n := len(*i64)
	c := cap(*i64)
	if n+1 > c {
		ns := make(uint64Slice, n, c*2)
		copy(ns, *i64)
		*i64 = ns
	}
	*i64 = (*i64)[0 : n+1]
	copy((*i64)[1:n+1], (*i64)[0:n])
	(*i64)[0] = v
	return true, n + 1
}

func (i64 *uint64Slice) PopBack() (bool, interface{}) {
	if i64 == nil {
		return false, nil
	}
	n := len(*i64)
	c := cap(*i64)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uint64Slice, n, c/2)
		copy(ns, *i64)
		*i64 = ns
	}
	x := (*i64)[n-1]
	*i64 = (*i64)[0 : n-1]
	return true, x
}

func (i64 *uint64Slice) PopFront() (bool, interface{}) {
	if i64 == nil {
		return false, nil
	}
	n := len(*i64)
	c := cap(*i64)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uint64Slice, n, c/2)
		copy(ns, *i64)
		*i64 = ns
	}
	x := (*i64)[0]
	*i64 = (*i64)[1:n]
	return true, x
}

func (i64 *uint64Slice) Index(x interface{}) (index int) {
	index = -1
	if i64 == nil {
		return
	}
	v, ok := x.(uint64)
	if !ok {
		return
	}

	for i := range *i64 {
		if (*i64)[i] == v {
			index = i
			break
		}
	}
	return
}

func (i64 *uint64Slice) IndexValue(index int) interface{} {
	if i64 == nil {
		return nil
	}
	n := len(*i64)
	if index < 0 || index >= n {
		return nil
	}
	return (*i64)[index]
}

func (i64 *uint64Slice) Range(fn func(index int, ele interface{}) bool) {
	if i64 == nil {
		return
	}
	for i, x := range *i64 {
		if fn(i, x) {
			break
		}
	}
}

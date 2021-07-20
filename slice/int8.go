package slice

import "fmt"

type int8Slice []int8

func newInt8Slice(c int) *int8Slice {
	s := make(int8Slice, 0, c)
	return &s
}

func (i8 *int8Slice) String() string {
	if i8 == nil {
		return ""
	}
	return fmt.Sprintf("%v", *i8)
}

func (i8 *int8Slice) PushBack(x interface{}) (bool, int) {
	if i8 == nil {
		return false, 0
	}
	v, ok := x.(int8)
	if !ok {
		return false, 0
	}
	n := len(*i8)
	c := cap(*i8)
	if n+1 > c {
		ns := make(int8Slice, n, c*2)
		copy(ns, *i8)
		*i8 = ns
	}
	*i8 = (*i8)[0 : n+1]
	(*i8)[n] = v
	return true, n + 1
}

func (i8 *int8Slice) PushFront(x interface{}) (bool, int) {
	if i8 == nil {
		return false, 0
	}
	v, ok := x.(int8)
	if !ok {
		return false, 0
	}
	n := len(*i8)
	c := cap(*i8)
	if n+1 > c {
		ns := make(int8Slice, n, c*2)
		copy(ns, *i8)
		*i8 = ns
	}
	*i8 = (*i8)[0 : n+1]
	copy((*i8)[1:n+1], (*i8)[0:n])
	(*i8)[0] = v
	return true, n + 1
}

func (i8 *int8Slice) PopBack() (bool, interface{}) {
	if i8 == nil {
		return false, nil
	}
	n := len(*i8)
	c := cap(*i8)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(int8Slice, n, c/2)
		copy(ns, *i8)
		*i8 = ns
	}
	x := (*i8)[n-1]
	*i8 = (*i8)[0 : n-1]
	return true, x
}

func (i8 *int8Slice) PopFront() (bool, interface{}) {
	if i8 == nil {
		return false, nil
	}
	n := len(*i8)
	c := cap(*i8)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(int8Slice, n, c/2)
		copy(ns, *i8)
		*i8 = ns
	}
	x := (*i8)[0]
	*i8 = (*i8)[1:n]
	return true, x
}

func (i8 *int8Slice) Index(x interface{}) (index int) {
	index = -1
	if i8 == nil {
		return
	}
	v, ok := x.(int8)
	if !ok {
		return
	}

	for i := range *i8 {
		if (*i8)[i] == v {
			index = i
			break
		}
	}
	return
}

func (i8 *int8Slice) IndexValue(index int) interface{} {
	if i8 == nil {
		return nil
	}
	n := len(*i8)
	if index < 0 || index >= n {
		return nil
	}
	return (*i8)[index]
}

func (i8 *int8Slice) Range(fn func(index int, ele interface{}) bool) {
	if i8 == nil {
		return
	}
	for i, x := range *i8 {
		if fn(i, x) {
			break
		}
	}
}

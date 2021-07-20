package slice

import "fmt"

type uint16Slice []uint16

func newUint16Slice(c int) *uint16Slice {
	s := make(uint16Slice, 0, c)
	return &s
}

func (u16 *uint16Slice) String() string {
	if u16 == nil {
		return ""
	}
	return fmt.Sprintf("%v", *u16)
}

func (u16 *uint16Slice) PushBack(x interface{}) (bool, int) {
	if u16 == nil {
		return false, 0
	}
	v, ok := x.(uint16)
	if !ok {
		return false, 0
	}
	n := len(*u16)
	c := cap(*u16)
	if n+1 > c {
		ns := make(uint16Slice, n, c*2)
		copy(ns, *u16)
		*u16 = ns
	}
	*u16 = (*u16)[0 : n+1]
	(*u16)[n] = v
	return true, n + 1
}

func (u16 *uint16Slice) PushFront(x interface{}) (bool, int) {
	if u16 == nil {
		return false, 0
	}
	v, ok := x.(uint16)
	if !ok {
		return false, 0
	}
	n := len(*u16)
	c := cap(*u16)
	if n+1 > c {
		ns := make(uint16Slice, n, c*2)
		copy(ns, *u16)
		*u16 = ns
	}
	*u16 = (*u16)[0 : n+1]
	copy((*u16)[1:n+1], (*u16)[0:n])
	(*u16)[0] = v
	return true, n + 1
}

func (u16 *uint16Slice) PopBack() (bool, interface{}) {
	if u16 == nil {
		return false, nil
	}
	n := len(*u16)
	c := cap(*u16)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uint16Slice, n, c/2)
		copy(ns, *u16)
		*u16 = ns
	}
	x := (*u16)[n-1]
	*u16 = (*u16)[0 : n-1]
	return true, x
}

func (u16 *uint16Slice) PopFront() (bool, interface{}) {
	if u16 == nil {
		return false, nil
	}
	n := len(*u16)
	c := cap(*u16)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uint16Slice, n, c/2)
		copy(ns, *u16)
		*u16 = ns
	}
	x := (*u16)[0]
	*u16 = (*u16)[1:n]
	return true, x
}

func (u16 *uint16Slice) Index(x interface{}) (index int) {
	index = -1
	if u16 == nil {
		return
	}
	v, ok := x.(uint16)
	if !ok {
		return
	}

	for i := range *u16 {
		if (*u16)[i] == v {
			index = i
			break
		}
	}
	return
}

func (u16 *uint16Slice) IndexValue(index int) interface{} {
	if u16 == nil {
		return nil
	}
	n := len(*u16)
	if index < 0 || index >= n {
		return nil
	}
	return (*u16)[index]
}

func (u16 *uint16Slice) Range(fn func(index int, ele interface{}) bool) {
	if u16 == nil {
		return
	}
	for i, x := range *u16 {
		if fn(i, x) {
			break
		}
	}
}

package slice

import (
	"fmt"
)

type intSlice []int

func newIntSlice(c int) *intSlice {
	s := make(intSlice, 0, c)
	return &s
}

func (is *intSlice) String() string {
	if is == nil {
		return ""
	}
	return fmt.Sprintf("%v", *is)
}

func (is *intSlice) PushBack(x interface{}) (bool, int) {
	if is == nil {
		return false, 0
	}
	v, ok := x.(int)
	if !ok {
		return false, 0
	}
	n := len(*is)
	c := cap(*is)
	if n+1 > c {
		ns := make(intSlice, n, c*2)
		copy(ns, *is)
		*is = ns
	}
	*is = (*is)[0 : n+1]
	(*is)[n] = v
	return true, n + 1
}

func (is *intSlice) PushFront(x interface{}) (bool, int) {
	if is == nil {
		return false, 0
	}
	v, ok := x.(int)
	if !ok {
		return false, 0
	}
	n := len(*is)
	c := cap(*is)
	if n+1 > c {
		ns := make(intSlice, n, c*2)
		copy(ns, *is)
		*is = ns
	}
	*is = (*is)[0 : n+1]
	copy((*is)[1:n+1], (*is)[0:n])
	(*is)[0] = v
	return true, n + 1
}

func (is *intSlice) PopBack() (bool, interface{}) {
	if is == nil {
		return false, nil
	}
	n := len(*is)
	c := cap(*is)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(intSlice, n, c/2)
		copy(ns, *is)
		*is = ns
	}
	x := (*is)[n-1]
	*is = (*is)[0 : n-1]
	return true, x
}

func (is *intSlice) PopFront() (bool, interface{}) {
	if is == nil {
		return false, nil
	}
	n := len(*is)
	c := cap(*is)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(intSlice, n, c/2)
		copy(ns, *is)
		*is = ns
	}
	x := (*is)[0]
	*is = (*is)[1:n]
	return true, x
}

func (is *intSlice) Index(x interface{}) (index int) {
	index = -1
	if is == nil {
		return
	}
	v, ok := x.(int)
	if !ok {
		return
	}

	for i := range *is {
		if (*is)[i] == v {
			index = i
			break
		}
	}
	return
}

func (is *intSlice) IndexValue(index int) interface{} {
	if is == nil {
		return nil
	}
	n := len(*is)
	if index < 0 || index >= n {
		return nil
	}
	return (*is)[index]
}

func (is *intSlice) Range(fn func(index int, ele interface{}) bool) {
	if is == nil {
		return
	}
	for i, x := range *is {
		if fn(i, x) {
			break
		}
	}
}

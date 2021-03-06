package slices

import (
	"fmt"

	"github.com/pyihe/go-pkg/sorts"
)

type int8Slice []int8

func newInt8Slice(c int) *int8Slice {
	s := make(int8Slice, 0, c)
	return &s
}

func (i8 *int8Slice) String() string {
	if i8 == nil {
		return ""
	}
	return fmt.Sprintf("len:[%v] cap:[%v] value:%v", len(*i8), cap(*i8), *i8)
}

func (i8 *int8Slice) Len() int {
	if i8 == nil {
		return 0
	}
	return len(*i8)
}

func (i8 *int8Slice) Cap() int {
	if i8 == nil {
		return 0
	}
	return cap(*i8)
}

func (i8 *int8Slice) Sort() {
	if i8 == nil {
		return
	}
	sorts.SortInt8s(*i8)
}

func (i8 *int8Slice) PushBack(x interface{}) (bool, int) {
	if i8 == nil {
		return false, 0
	}
	ok, v := convertToInt8(x)
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
	ok, v := convertToInt8(x)
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
	ok, v := convertToInt8(x)
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

func (i8 *int8Slice) Delete(x interface{}) (ok bool) {
	if i8 == nil {
		return
	}
	if len(*i8) == 0 {
		return
	}
	ok, v := convertToInt8(x)
	if !ok {
		return
	}
	for i, e := range *i8 {
		if e == v {
			*i8 = append((*i8)[:i], (*i8)[i+1:]...)
			ok = true
			break
		}
	}
	return
}

package slice

import (
	"fmt"

	"github.com/pyihe/go-pkg/sorts"
)

type int16Slice []int16

func newInt16Slice(c int) *int16Slice {
	s := make(int16Slice, 0, c)
	return &s
}

func (i16 *int16Slice) String() string {
	if i16 == nil {
		return ""
	}
	return fmt.Sprintf("len:[%v] cap:[%v] value:%v", len(*i16), cap(*i16), *i16)
}

func (i16 *int16Slice) Len() int {
	if i16 == nil {
		return 0
	}
	return len(*i16)
}

func (i16 *int16Slice) Cap() int {
	if i16 == nil {
		return 0
	}
	return cap(*i16)
}

func (i16 *int16Slice) Sort() {
	if i16 == nil {
		return
	}
	sorts.SortInt16s(*i16)
}

func (i16 *int16Slice) PushBack(x interface{}) (bool, int) {
	if i16 == nil {
		return false, 0
	}
	ok, v := convertToInt16(x)
	if !ok {
		return false, 0
	}
	n := len(*i16)
	c := cap(*i16)
	if n+1 > c {
		ns := make(int16Slice, n, c*2)
		copy(ns, *i16)
		*i16 = ns
	}
	*i16 = (*i16)[0 : n+1]
	(*i16)[n] = v
	return true, n + 1
}

func (i16 *int16Slice) PushFront(x interface{}) (bool, int) {
	if i16 == nil {
		return false, 0
	}
	ok, v := convertToInt16(x)
	if !ok {
		return false, 0
	}
	n := len(*i16)
	c := cap(*i16)
	if n+1 > c {
		ns := make(int16Slice, n, c*2)
		copy(ns, *i16)
		*i16 = ns
	}
	*i16 = (*i16)[0 : n+1]
	copy((*i16)[1:n+1], (*i16)[0:n])
	(*i16)[0] = v
	return true, n + 1
}

func (i16 *int16Slice) PopBack() (bool, interface{}) {
	if i16 == nil {
		return false, nil
	}
	n := len(*i16)
	c := cap(*i16)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(int16Slice, n, c/2)
		copy(ns, *i16)
		*i16 = ns
	}
	x := (*i16)[n-1]
	*i16 = (*i16)[0 : n-1]
	return true, x
}

func (i16 *int16Slice) PopFront() (bool, interface{}) {
	if i16 == nil {
		return false, nil
	}
	n := len(*i16)
	c := cap(*i16)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(int16Slice, n, c/2)
		copy(ns, *i16)
		*i16 = ns
	}
	x := (*i16)[0]
	*i16 = (*i16)[1:n]
	return true, x
}

func (i16 *int16Slice) Index(x interface{}) (index int) {
	index = -1
	if i16 == nil {
		return
	}
	ok, v := convertToInt16(x)
	if !ok {
		return
	}

	for i := range *i16 {
		if (*i16)[i] == v {
			index = i
			break
		}
	}
	return
}

func (i16 *int16Slice) IndexValue(index int) interface{} {
	if i16 == nil {
		return nil
	}
	n := len(*i16)
	if index < 0 || index >= n {
		return nil
	}
	return (*i16)[index]
}

func (i16 *int16Slice) Range(fn func(index int, ele interface{}) bool) {
	if i16 == nil {
		return
	}
	for i, x := range *i16 {
		if fn(i, x) {
			break
		}
	}
}

func (i16 *int16Slice) Delete(x interface{}) (ok bool) {
	if i16 == nil {
		return
	}
	if len(*i16) == 0 {
		return
	}
	ok, v := convertToInt16(x)
	if !ok {
		return
	}
	for i, e := range *i16 {
		if e == v {
			*i16 = append((*i16)[:i], (*i16)[i+1:]...)
			ok = true
			break
		}
	}
	return
}

package slicepkg

import (
	"fmt"

	"github.com/pyihe/go-pkg/sortpkg"
)

type int32Slice []int32

func newInt32Slice(c int) *int32Slice {
	s := make(int32Slice, 0, c)
	return &s
}

func (i32 *int32Slice) String() string {
	if i32 == nil {
		return ""
	}
	return fmt.Sprintf("len:[%v] cap:[%v] value:%v", len(*i32), cap(*i32), *i32)
}

func (i32 *int32Slice) Len() int {
	if i32 == nil {
		return 0
	}
	return len(*i32)
}

func (i32 *int32Slice) Cap() int {
	if i32 == nil {
		return 0
	}
	return cap(*i32)
}

func (i32 *int32Slice) Sort() {
	if i32 == nil {
		return
	}
	sortpkg.SortInt32s(*i32)
}

func (i32 *int32Slice) PushBack(x interface{}) (bool, int) {
	if i32 == nil {
		return false, 0
	}
	ok, v := convertToInt32(x)
	if !ok {
		return false, 0
	}
	n := len(*i32)
	c := cap(*i32)
	if n+1 > c {
		ns := make(int32Slice, n, c*2)
		copy(ns, *i32)
		*i32 = ns
	}
	*i32 = (*i32)[0 : n+1]
	(*i32)[n] = v
	return true, n + 1
}

func (i32 *int32Slice) PushFront(x interface{}) (bool, int) {
	if i32 == nil {
		return false, 0
	}
	ok, v := convertToInt32(x)
	if !ok {
		return false, 0
	}
	n := len(*i32)
	c := cap(*i32)
	if n+1 > c {
		ns := make(int32Slice, n, c*2)
		copy(ns, *i32)
		*i32 = ns
	}
	*i32 = (*i32)[0 : n+1]
	copy((*i32)[1:n+1], (*i32)[0:n])
	(*i32)[0] = v
	return true, n + 1
}

func (i32 *int32Slice) PopBack() (bool, interface{}) {
	if i32 == nil {
		return false, nil
	}
	n := len(*i32)
	c := cap(*i32)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(int32Slice, n, c/2)
		copy(ns, *i32)
		*i32 = ns
	}
	x := (*i32)[n-1]
	*i32 = (*i32)[0 : n-1]
	return true, x
}

func (i32 *int32Slice) PopFront() (bool, interface{}) {
	if i32 == nil {
		return false, nil
	}
	n := len(*i32)
	c := cap(*i32)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(int32Slice, n, c/2)
		copy(ns, *i32)
		*i32 = ns
	}
	x := (*i32)[0]
	*i32 = (*i32)[1:n]
	return true, x
}

func (i32 *int32Slice) Index(x interface{}) (index int) {
	index = -1
	if i32 == nil {
		return
	}
	ok, v := convertToInt32(x)
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

func (i32 *int32Slice) IndexValue(index int) interface{} {
	if i32 == nil {
		return nil
	}
	n := len(*i32)
	if index < 0 || index >= n {
		return nil
	}
	return (*i32)[index]
}

func (i32 *int32Slice) Range(fn func(index int, ele interface{}) bool) {
	if i32 == nil {
		return
	}
	for i, x := range *i32 {
		if fn(i, x) {
			break
		}
	}
}

func (i32 *int32Slice) Delete(x interface{}) (ok bool) {
	if i32 == nil {
		return
	}
	if len(*i32) == 0 {
		return
	}
	ok, v := convertToInt32(x)
	if !ok {
		return
	}
	for i, e := range *i32 {
		if e == v {
			*i32 = append((*i32)[:i], (*i32)[i+1:]...)
			ok = true
			break
		}
	}
	return
}

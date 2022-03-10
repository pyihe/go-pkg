package slices

import (
	"fmt"

	"github.com/pyihe/go-pkg/sorts"
)

type float64Slice []float64

func newFloat64Slice(c int) *float64Slice {
	s := make(float64Slice, 0, c)
	return &s
}

func (f *float64Slice) String() string {
	if f == nil {
		return ""
	}
	return fmt.Sprintf("len:[%v] cap:[%v] value:%v", len(*f), cap(*f), *f)
}

func (f *float64Slice) Len() int {
	if f == nil {
		return 0
	}
	return len(*f)
}

func (f *float64Slice) Cap() int {
	if f == nil {
		return 0
	}
	return cap(*f)
}

func (f *float64Slice) Sort() {
	if f == nil {
		return
	}
	sorts.SortFloat64s(*f)
}

func (f *float64Slice) PushBack(x interface{}) (bool, int) {
	if f == nil {
		return false, 0
	}
	ok, v := convertToFloat64(x)
	if !ok {
		return false, 0
	}
	n := len(*f)
	c := cap(*f)
	if n+1 > c {
		ns := make(float64Slice, n, c*2)
		copy(ns, *f)
		*f = ns
	}
	*f = (*f)[0 : n+1]
	(*f)[n] = v
	return true, n + 1
}

func (f *float64Slice) PushFront(x interface{}) (bool, int) {
	if f == nil {
		return false, 0
	}
	ok, v := convertToFloat64(x)
	if !ok {
		return false, 0
	}
	n := len(*f)
	c := cap(*f)
	if n+1 > c {
		ns := make(float64Slice, n, c*2)
		copy(ns, *f)
		*f = ns
	}
	*f = (*f)[0 : n+1]
	copy((*f)[1:n+1], (*f)[0:n])
	(*f)[0] = v
	return true, n + 1
}

func (f *float64Slice) PopBack() (bool, interface{}) {
	if f == nil {
		return false, nil
	}
	n := len(*f)
	c := cap(*f)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(float64Slice, n, c/2)
		copy(ns, *f)
		*f = ns
	}
	x := (*f)[n-1]
	*f = (*f)[0 : n-1]
	return true, x
}

func (f *float64Slice) PopFront() (bool, interface{}) {
	if f == nil {
		return false, nil
	}
	n := len(*f)
	c := cap(*f)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(float64Slice, n, c/2)
		copy(ns, *f)
		*f = ns
	}
	x := (*f)[0]
	*f = (*f)[1:n]
	return true, x
}

func (f *float64Slice) Index(x interface{}) (index int) {
	index = -1
	if f == nil {
		return
	}
	ok, v := convertToFloat64(x)
	if !ok {
		return
	}

	for i := range *f {
		if (*f)[i] == v {
			index = i
			break
		}
	}
	return
}

func (f *float64Slice) IndexValue(index int) interface{} {
	if f == nil {
		return nil
	}
	n := len(*f)
	if index < 0 || index >= n {
		return nil
	}
	return (*f)[index]
}

func (f *float64Slice) Range(fn func(index int, ele interface{}) bool) {
	if f == nil {
		return
	}
	for i, x := range *f {
		if fn(i, x) {
			break
		}
	}
}

func (f *float64Slice) Delete(x interface{}) (ok bool) {
	if f == nil {
		return
	}
	if len(*f) == 0 {
		return
	}
	ok, v := convertToFloat64(x)
	if !ok {
		return false
	}
	for i, e := range *f {
		if e == v {
			*f = append((*f)[:i], (*f)[i+1:]...)
			ok = true
			break
		}
	}
	return
}

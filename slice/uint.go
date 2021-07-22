package slice

import (
	"fmt"

	"github.com/pyihe/go-pkg/sorts"
)

type uintSlice []uint

func newUintSlice(c int) *uintSlice {
	s := make(uintSlice, 0, c)
	return &s
}

func (us *uintSlice) String() string {
	if us == nil {
		return ""
	}
	return fmt.Sprintf("len:[%v] cap:[%v] value:%v", len(*us), cap(*us), *us)
}

func (us *uintSlice) Len() int {
	if us == nil {
		return 0
	}
	return len(*us)
}

func (us *uintSlice) Cap() int {
	if us == nil {
		return 0
	}
	return cap(*us)
}

func (us *uintSlice) Sort() {
	if us == nil {
		return
	}
	sorts.SortUints(*us)
}

func (us *uintSlice) PushBack(x interface{}) (bool, int) {
	if us == nil {
		return false, 0
	}
	ok, v := convertToUint(x)
	if !ok {
		return false, 0
	}
	n := len(*us)
	c := cap(*us)
	if n+1 > c {
		ns := make(uintSlice, n, c*2)
		copy(ns, *us)
		*us = ns
	}
	*us = (*us)[0 : n+1]
	(*us)[n] = v
	return true, n + 1
}

func (us *uintSlice) PushFront(x interface{}) (bool, int) {
	if us == nil {
		return false, 0
	}
	ok, v := convertToUint(x)
	if !ok {
		return false, 0
	}
	n := len(*us)
	c := cap(*us)
	if n+1 > c {
		ns := make(uintSlice, n, c*2)
		copy(ns, *us)
		*us = ns
	}
	*us = (*us)[0 : n+1]
	copy((*us)[1:n+1], (*us)[0:n])
	(*us)[0] = v
	return true, n + 1
}

func (us *uintSlice) PopBack() (bool, interface{}) {
	if us == nil {
		return false, nil
	}
	n := len(*us)
	c := cap(*us)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uintSlice, n, c/2)
		copy(ns, *us)
		*us = ns
	}
	x := (*us)[n-1]
	*us = (*us)[0 : n-1]
	return true, x
}

func (us *uintSlice) PopFront() (bool, interface{}) {
	if us == nil {
		return false, nil
	}
	n := len(*us)
	c := cap(*us)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(uintSlice, n, c/2)
		copy(ns, *us)
		*us = ns
	}
	x := (*us)[0]
	*us = (*us)[1:n]
	return true, x
}

func (us *uintSlice) Index(x interface{}) (index int) {
	index = -1
	if us == nil {
		return
	}
	ok, v := convertToUint(x)
	if !ok {
		return
	}

	for i := range *us {
		if (*us)[i] == v {
			index = i
			break
		}
	}
	return
}

func (us *uintSlice) IndexValue(index int) interface{} {
	if us == nil {
		return nil
	}
	n := len(*us)
	if index < 0 || index >= n {
		return nil
	}
	return (*us)[index]
}

func (us *uintSlice) Range(fn func(index int, ele interface{}) bool) {
	if us == nil {
		return
	}
	for i, x := range *us {
		if fn(i, x) {
			break
		}
	}
}

func (us *uintSlice) Delete(x interface{}) (ok bool) {
	if us == nil {
		return
	}
	if len(*us) == 0 {
		return
	}
	ok, v := convertToUint(x)
	if !ok {
		return
	}
	for i, e := range *us {
		if e == v {
			*us = append((*us)[:i], (*us)[i+1:]...)
			ok = true
			break
		}
	}
	return
}

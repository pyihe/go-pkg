package slice

import (
	"fmt"

	"github.com/pyihe/go-pkg/sorts"
)

type stringSlice []string

func newStringSlice(c int) *stringSlice {
	s := make(stringSlice, 0, c)
	return &s
}

func (ss *stringSlice) String() string {
	if ss == nil {
		return ""
	}
	return fmt.Sprintf("len:[%v] cap:[%v] value:%v", len(*ss), cap(*ss), *ss)
}

func (ss *stringSlice) Len() int {
	if ss == nil {
		return 0
	}
	return len(*ss)
}

func (ss *stringSlice) Cap() int {
	if ss == nil {
		return 0
	}
	return cap(*ss)
}

func (ss *stringSlice) Sort() {
	if ss == nil {
		return
	}
	sorts.SortStrings(*ss)
}

func (ss *stringSlice) PushBack(x interface{}) (bool, int) {
	if ss == nil {
		return false, 0
	}
	ok, v := convertToIntString(x)
	if !ok {
		return false, 0
	}
	n := len(*ss)
	c := cap(*ss)
	if n+1 > c {
		ns := make(stringSlice, n, c*2)
		copy(ns, *ss)
		*ss = ns
	}
	*ss = (*ss)[0 : n+1]
	(*ss)[n] = v
	return true, n + 1
}

func (ss *stringSlice) PushFront(x interface{}) (bool, int) {
	if ss == nil {
		return false, 0
	}
	ok, v := convertToIntString(x)
	if !ok {
		return false, 0
	}
	n := len(*ss)
	c := cap(*ss)
	if n+1 > c {
		ns := make(stringSlice, n, c*2)
		copy(ns, *ss)
		*ss = ns
	}
	*ss = (*ss)[0 : n+1]
	copy((*ss)[1:n+1], (*ss)[0:n])
	(*ss)[0] = v
	return true, n + 1
}

func (ss *stringSlice) PopBack() (bool, interface{}) {
	if ss == nil {
		return false, nil
	}
	n := len(*ss)
	c := cap(*ss)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(stringSlice, n, c/2)
		copy(ns, *ss)
		*ss = ns
	}
	x := (*ss)[n-1]
	*ss = (*ss)[0 : n-1]
	return true, x
}

func (ss *stringSlice) PopFront() (bool, interface{}) {
	if ss == nil {
		return false, nil
	}
	n := len(*ss)
	c := cap(*ss)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(stringSlice, n, c/2)
		copy(ns, *ss)
		*ss = ns
	}
	x := (*ss)[0]
	*ss = (*ss)[1:n]
	return true, x
}

func (ss *stringSlice) Index(x interface{}) (index int) {
	index = -1
	if ss == nil {
		return
	}
	ok, v := convertToIntString(x)
	if !ok {
		return
	}

	for i := range *ss {
		if (*ss)[i] == v {
			index = i
			break
		}
	}
	return
}

func (ss *stringSlice) IndexValue(index int) interface{} {
	if ss == nil {
		return nil
	}
	n := len(*ss)
	if index < 0 || index >= n {
		return nil
	}
	return (*ss)[index]
}

func (ss *stringSlice) Range(fn func(index int, ele interface{}) bool) {
	if ss == nil {
		return
	}
	for i, x := range *ss {
		if fn(i, x) {
			break
		}
	}
}

func (ss *stringSlice) Delete(x interface{}) (ok bool) {
	if ss == nil {
		return
	}
	if len(*ss) == 0 {
		return
	}
	ok, v := convertToIntString(x)
	if !ok {
		return
	}
	for i, e := range *ss {
		if e == v {
			*ss = append((*ss)[:i], (*ss)[i+1:]...)
			ok = true
			break
		}
	}
	return
}

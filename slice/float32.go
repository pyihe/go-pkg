package slice

import "fmt"

type float32Slice []float32

func newFloat32Slice(c int) *float32Slice {
	s := make(float32Slice, 0, c)
	return &s
}

func (f *float32Slice) String() string {
	if f == nil {
		return ""
	}
	return fmt.Sprintf("%v", *f)
}

func (f *float32Slice) PushBack(x interface{}) (bool, int) {
	if f == nil {
		return false, 0
	}
	v, ok := x.(float32)
	if !ok {
		return false, 0
	}
	n := len(*f)
	c := cap(*f)
	if n+1 > c {
		ns := make(float32Slice, n, c*2)
		copy(ns, *f)
		*f = ns
	}
	*f = (*f)[0 : n+1]
	(*f)[n] = v
	return true, n + 1
}

func (f *float32Slice) PushFront(x interface{}) (bool, int) {
	if f == nil {
		return false, 0
	}
	v, ok := x.(float32)
	if !ok {
		return false, 0
	}
	n := len(*f)
	c := cap(*f)
	if n+1 > c {
		ns := make(float32Slice, n, c*2)
		copy(ns, *f)
		*f = ns
	}
	*f = (*f)[0 : n+1]
	copy((*f)[1:n+1], (*f)[0:n])
	(*f)[0] = v
	return true, n + 1
}

func (f *float32Slice) PopBack() (bool, interface{}) {
	if f == nil {
		return false, nil
	}
	n := len(*f)
	c := cap(*f)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(float32Slice, n, c/2)
		copy(ns, *f)
		*f = ns
	}
	x := (*f)[n-1]
	*f = (*f)[0 : n-1]
	return true, x
}

func (f *float32Slice) PopFront() (bool, interface{}) {
	if f == nil {
		return false, nil
	}
	n := len(*f)
	c := cap(*f)
	if n == 0 {
		return false, nil
	}
	if n < (c/2) && c > 25 {
		ns := make(float32Slice, n, c/2)
		copy(ns, *f)
		*f = ns
	}
	x := (*f)[0]
	*f = (*f)[1:n]
	return true, x
}

func (f *float32Slice) Index(x interface{}) (index int) {
	index = -1
	if f == nil {
		return
	}
	v, ok := x.(float32)
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

func (f *float32Slice) IndexValue(index int) interface{} {
	if f == nil {
		return nil
	}
	n := len(*f)
	if index < 0 || index >= n {
		return nil
	}
	return (*f)[index]
}

func (f *float32Slice) Range(fn func(index int, ele interface{}) bool) {
	if f == nil {
		return
	}
	for i, x := range *f {
		if fn(i, x) {
			break
		}
	}
}

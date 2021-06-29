package queue

import (
	"sync"
)

type (
	Queue interface {
		Len() int
		UnsafeLen() int
		Push(ele interface{})
		Pop() interface{}
		Del(i int)
		UnsafeDel(i int)
		Get(i int) interface{}
		UnsafeGet(i int) interface{}
		Set(index int, v interface{})
		UnsafeSet(index int, v interface{})
		Range(func(i int, v interface{}))
		UnsafeRange(func(i int, v interface{}))
		Index(data interface{}) (ok bool, i int)
		UnsafeIndex(data interface{}) (ok bool, i int)
	}

	queue struct {
		mu    *sync.RWMutex
		t     ListType
		count int
		data  []interface{}
	}

	ListType int
)

const (
	ListTypeQueue ListType = iota + 1
	ListTypeStack
)

var _ Queue = &queue{}

func NewQueue(t ListType, defaultCap int) Queue {
	switch t {
	case ListTypeStack:
	case ListTypeQueue:
	default:
		panic("unknown list type")
	}
	q := &queue{
		mu:    &sync.RWMutex{},
		count: 0,
		t:     t,
		data:  make([]interface{}, 0, defaultCap),
	}
	return q
}

func (q *queue) init() {
	q.count = 0
	q.data = make([]interface{}, 0)
}

func (q *queue) checkLen(i int) {
	if q.count-1 < i {
		panic("out of range")
	}
}

func (q *queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.count
}

func (q *queue) UnsafeLen() int {
	return q.count
}

//add
func (q *queue) Push(ele interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.count++
	q.data = append(q.data, ele)
}

//get&remove
func (q *queue) Pop() interface{} {
	var data interface{}
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.count == 0 {
		return nil
	}
	switch q.t {
	case ListTypeQueue:
		data = q.data[0]
		q.data = q.data[1:]
	case ListTypeStack:
		data = q.data[q.count-1]
		q.data = q.data[:q.count-1]
	}
	q.count--
	return data
}

//update
func (q *queue) Set(i int, v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.checkLen(i)
	q.data[i] = v
}

func (q *queue) UnsafeSet(i int, v interface{}) {
	q.checkLen(i)
	q.data[i] = v
}

//del
func (q *queue) Del(i int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.checkLen(i)
	q.data = append(q.data[:i], q.data[i+1:]...)
	q.count--
}

func (q *queue) UnsafeDel(i int) {
	q.checkLen(i)
	q.data = append(q.data[:i], q.data[i+1:]...)
	q.count--
}

//Get
func (q *queue) Get(i int) interface{} {
	q.mu.RLock()
	defer q.mu.RUnlock()
	q.checkLen(i)
	d := q.data[i]
	return d
}

func (q *queue) UnsafeGet(i int) interface{} {
	q.checkLen(i)
	d := q.data[i]
	return d
}

//safe range
func (q *queue) Range(f func(i int, v interface{})) {
	q.mu.Lock()
	defer q.mu.Unlock()
	switch q.t {
	case ListTypeQueue:
		for i, v := range q.data {
			f(i, v)
		}
	case ListTypeStack:
		for i := q.count - 1; i >= 0; i-- {
			f(i, q.data[i])
		}
	}

}

//unsafe range
func (q *queue) UnsafeRange(f func(i int, v interface{})) {
	switch q.t {
	case ListTypeQueue:
		for i, v := range q.data {
			f(i, v)
		}
	case ListTypeStack:
		for i := q.count - 1; i >= 0; i-- {
			f(i, q.data[i])
		}
	}
}

func (q *queue) Index(data interface{}) (bool, int) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for i, v := range q.data {
		if v == data {
			return true, i
		}
	}
	return false, 0
}

func (q *queue) UnsafeIndex(data interface{}) (bool, int) {
	for i, v := range q.data {
		if v == data {
			return true, i
		}
	}
	return false, 0
}

package queue

import "github.com/pyihe/go-pkg/errors"

const defaultQueueSize = 129

type linearQueue struct {
	data  []interface{} // 元素
	size  int           // 队列长度
	front int           // 队首
	rear  int           // 队尾
}

func newLinearQueue(size int) *linearQueue {
	if size <= 0 {
		size = defaultQueueSize
	} else {
		size += 1 // 因为循环队列采用的是队首与队尾差一个单元时为队列满, 所以加1
	}
	return &linearQueue{
		data:  make([]interface{}, size, size),
		size:  size,
		front: 0,
		rear:  0,
	}
}

func (q *linearQueue) full() bool {
	return (q.rear+1)%q.size == q.front
}

func (q *linearQueue) empty() bool {
	return q.front == q.rear
}

func (q *linearQueue) Reset() {
	*q = linearQueue{
		data:  make([]interface{}, q.size, q.size),
		size:  q.size,
		front: 0,
		rear:  0,
	}
}

func (q *linearQueue) Len() int {
	return (q.rear - q.front + q.size) % q.size
}

func (q *linearQueue) Head() (interface{}, bool) {
	if q.Len() == 0 {
		return nil, false
	}
	return q.data[q.front], true
}

func (q *linearQueue) EnQueue(ele interface{}) error {
	if q.full() {
		return errors.New("full queue")
	}
	q.data[q.rear] = ele
	q.rear = (q.rear + 1) % q.size
	return nil
}

func (q *linearQueue) DeQueue() (interface{}, error) {
	if q.empty() {
		return nil, errors.New("empty queue")
	}
	ele := q.data[q.front]
	q.front = (q.front + 1) % q.size
	return ele, nil
}

func (q *linearQueue) Each(handler func(ele interface{}) bool) {
	if handler == nil || q.empty() {
		return
	}
	p := q.front
	for p != q.rear {
		if handler(q.data[p]) {
			break
		}
		p += 1
		if p == q.size {
			p = 0
		}
	}
}

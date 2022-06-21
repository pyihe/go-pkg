package queue

import "github.com/pyihe/go-pkg/errors"

type node struct {
	data interface{}
	next *node
}

type chainQueue struct {
	front  *node // 队首指针, 指向链式队列的头节点, 头节点不存储数据
	rear   *node // 队尾指针, 指向链式队列的尾节点
	length int   // 队列长度, 如果不用length记录, 则每次获取长度时都需要对链表遍历, 时间复杂度为O(N)
}

func newChainQueue() *chainQueue {
	head := &node{
		data: nil,
		next: nil,
	}
	return &chainQueue{
		front:  head,
		rear:   head,
		length: 0,
	}
}

func (q *chainQueue) Reset() {
	*q = chainQueue{
		front:  nil,
		rear:   nil,
		length: 0,
	}
}

func (q *chainQueue) Len() int {
	return q.length
}

func (q *chainQueue) Head() (interface{}, bool) {
	if q.front == q.rear {
		return nil, false
	}
	return q.front.next.data, true
}

func (q *chainQueue) EnQueue(ele interface{}) error {
	ne := &node{
		data: ele,
		next: nil,
	}
	q.rear.next = ne
	q.rear = ne
	q.length += 1
	return nil
}

func (q *chainQueue) DeQueue() (interface{}, error) {
	if q.front == q.rear {
		return nil, errors.New("empty queue")
	}
	p := q.front.next // 第一个节点
	ele := p.data
	q.front.next = p.next
	q.length -= 1
	// 判断是否出队列后没有元素了, 即队尾和第一个节点相同, 此时需要将队尾指针指向队首指针
	if q.rear == p {
		q.rear = q.front
	}
	return ele, nil
}

func (q *chainQueue) Each(handler func(ele interface{}) bool) {
	if handler == nil {
		return
	}
	p := q.front.next
	for p != nil {
		if handler(p.data) {
			break
		}
		p = p.next
	}
}

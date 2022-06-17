package list

import "fmt"

type node struct {
	data interface{}
	next *node
}

// 单链表线性表
type singleList struct {
	length int   // 链表长度
	head   *node // 头节点
}

func newSingleList() *singleList {
	return &singleList{}
}

func (sl *singleList) String() string {
	result := ""
	p := sl.head
	for p != nil {
		result += fmt.Sprintf("%v-", p.data)
		p = p.next
	}
	return result
}

func (sl *singleList) IsEmpty() bool {
	return sl.length == 0
}

func (sl *singleList) Len() int {
	return sl.length
}

func (sl *singleList) Reset() {
	*sl = singleList{
		length: 0,
		head:   nil,
	}
}

func (sl *singleList) Get(pos int) (interface{}, bool) {
	if sl.length == 0 || pos <= 0 || pos > sl.length {
		return nil, false
	}
	num := 1
	p := sl.head
	for p != nil && num < pos {
		p = p.next
		num += 1
	}
	if p != nil && num == pos {
		return p.data, true
	}
	return nil, false
}

func (sl *singleList) Pos(ele interface{}) int {
	num := 1
	p := sl.head
	for p != nil && p.data != ele {
		p = p.next
		num += 1
	}
	if p != nil {
		return num
	}
	return 0
}

func (sl *singleList) Insert(pos int, ele interface{}) error {
	if pos > 0 && pos <= sl.length+1 {
		if sl.head == nil && pos == 1 {
			sl.head = &node{
				data: ele,
				next: nil,
			}
			sl.length += 1
			return nil
		}
		num := 0
		p := sl.head
		for p != nil && num+1 < pos {
			p = p.next
			num += 1
		}
		if p != nil && num+1 == pos {
			s := &node{} // 新节点
			s.data = ele
			s.next = p.next
			p.next = s
			sl.length += 1
			return nil
		}
	}
	return ErrInvalidPos
}

func (sl *singleList) Delete(pos int) bool {
	if pos > 0 && pos <= sl.length {
		num := 1
		p := sl.head
		for p != nil && num+1 < pos {
			p = p.next
			num += 1
		}
		if p != nil {
			switch {
			case pos == 1:
				sl.head.data = nil
				sl.head = sl.head.next
			default:
				if p.next != nil {
					p.next.data = nil
					p.next = p.next.next
				}
			}
			sl.length -= 1
			return true
		}
	}
	return false
}

package list

type circleList struct {
	length int
	head   *node
}

func (cl *circleList) IsEmpty() bool {
	return cl.length == 0
}

func (cl *circleList) Len() int {
	return cl.length
}

func (cl *circleList) Reset() {
	*cl = circleList{
		length: 0,
		head:   nil,
	}
}

func (cl *circleList) Get(pos int) (interface{}, bool) {
	if cl.length == 0 || pos <= 0 || pos > cl.length {
		return nil, false
	}
	num := 1
	p := cl.head
	for p != nil && p.next != cl.head && num < pos {
		p = p.next
		num += 1
	}
	if p != nil && num == pos {
		return p.data, true
	}
	return nil, false
}

func (cl *circleList) Pos(ele interface{}) int {
	num := 1
	p := cl.head
	for p != nil && p.next != cl.head && p.data != ele {
		p = p.next
		num += 1
	}
	if p != nil {
		return num
	}
	return 0
}

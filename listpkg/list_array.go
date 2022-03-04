package listpkg

import "sync"

const (
	defaultCap = 128
)

// ArrayList 切片实现的队列
type ArrayList struct {
	mu   *sync.Mutex
	data []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		mu:   &sync.Mutex{},
		data: make([]interface{}, 0, defaultCap),
	}
}

// LPush 从队首添加元素
func (array *ArrayList) LPush(datas ...interface{}) {
	array.mu.Lock()
	defer array.mu.Unlock()

	n := len(datas)
	if n == 0 {
		return
	}

	length := len(array.data)
	capacity := cap(array.data)
	switch length == capacity {
	case true: // 如果已经放满了，需要扩容
		na := make([]interface{}, length+1, capacity+capacity/2)
		copy(na[1:], array.data)
		na[0] = datas
		array.data = na
	default: // 没有放满的话
		array.data = append(array.data, nil)
		for i := length; i > 0; i-- {
			array.data[i] = array.data[i-1]
		}
		array.data[0] = datas
	}
}

// LPop 从队首取数据
func (array *ArrayList) LPop() (data interface{}) {
	array.mu.Lock()
	defer array.mu.Unlock()
	if len(array.data) == 0 {
		return
	}
	data = array.data[0]
	array.data = array.data[1:]
	return
}

// RPush 从队尾添加元素
func (array *ArrayList) RPush(data interface{}) {
	array.mu.Lock()
	array.data = append(array.data, data)
	array.mu.Unlock()
}

// RPop 从队尾取出并删除元素
func (array *ArrayList) RPop() (data interface{}) {
	array.mu.Lock()
	defer array.mu.Unlock()

	n := len(array.data)
	if n == 0 {
		return
	}

	data = array.data[n-1]
	array.data = array.data[:n-1]
	return data
}

func (array *ArrayList) Data() []interface{} {
	return array.data
}

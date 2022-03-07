package listpkg

import (
	"sync"
)

// ArrayList 切片实现的队列
type ArrayList struct {
	mu   *sync.Mutex
	data []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		mu: &sync.Mutex{},
	}
}

// LPush 从队首添加元素
func (array *ArrayList) LPush(elements ...interface{}) (n int) {
	elementCount := len(elements)
	if elementCount == 0 {
		return
	}
	array.mu.Lock()
	array.data = append(elements, array.data...)
	n = elementCount
	array.mu.Unlock()
	return
}

// LPop 从队首取数据
func (array *ArrayList) LPop(n int) (data []interface{}) {
	array.mu.Lock()
	if len(array.data) > 0 {
		data, array.data = array.data[:n], array.data[n:]
	}
	array.mu.Unlock()
	return
}

// LInsert 在指定位置mark([0,len))前面插入数据data
func (array *ArrayList) LInsert(mark int, val ...interface{}) (n int) {
	vCount := len(val)
	if vCount == 0 {
		return
	}
	var insertFunc = func(s []interface{}, i int, vs ...interface{}) []interface{} {
		if cnt := len(s) + len(vs); cnt <= cap(s) {
			s2 := s[:cnt]
			copy(s2[i+len(vs):], s[i:])
			copy(s2[i:], vs)
			return s2
		}
		s2 := make([]interface{}, len(s)+len(vs))
		copy(s2, s[:i])
		copy(s2[i:], vs)
		copy(s2[i+len(vs):], s[i:])
		return s2
	}

	array.mu.Lock()
	defer array.mu.Unlock()

	cnt := len(array.data)
	// 如果队列为空，直接将val添加进队列
	if cnt == 0 {
		array.data = append(array.data, val...)
		n = vCount
		return
	}
	// 如果队列不为空，则在指定位置添加
	// 指定位置必须合法
	if mark < 0 || mark >= cnt {
		return
	}
	array.data = insertFunc(array.data, mark, val)
	n = vCount
	return
}

// RPush 从队尾添加元素
func (array *ArrayList) RPush(elements ...interface{}) (n int) {
	array.mu.Lock()
	array.data = append(array.data, elements...)
	n = len(elements)
	array.mu.Unlock()
	return
}

// RPop 从队尾取出并删除元素
func (array *ArrayList) RPop() (data interface{}) {
	array.mu.Lock()
	if n := len(array.data); n > 0 {
		data, array.data = array.data[n-1], array.data[:n-1]
	}
	array.mu.Unlock()
	return
}

// RInsert 在指定位置mark([0, len))后面添加数据
func (array *ArrayList) RInsert(mark int, val ...interface{}) (n int) {
	vCount := len(val)
	if vCount == 0 {
		return
	}
	var insertFunc = func(s []interface{}, i int, vs ...interface{}) []interface{} {
		// 如果是在最后一个元素的右边插入数据，直接append即可
		if i == len(s)-1 {
			return append(s, vs...)
		}
		// 不需要扩容
		if cnt := len(s) + len(vs); cnt <= cap(s) {
			s2 := s[:cnt]
			copy(s2[i+len(vs)+1:], s[i+1:])
			copy(s2[i+1:], vs)
			return s2
		}
		s2 := make([]interface{}, len(s)+len(vs))
		copy(s2, s[:i+1])
		copy(s2[i+1:], vs)
		copy(s2[i+len(vs)+1:], s[i+1:])
		return s2
	}

	array.mu.Lock()
	defer array.mu.Unlock()

	cnt := len(array.data)
	// 如果队列为空，直接将val添加进队列
	if cnt == 0 {
		array.data = append(array.data, val...)
		n = vCount
		return
	}
	// 如果队列不为空，则在指定位置添加
	// 指定位置必须合法
	if mark < 0 || mark >= cnt {
		return
	}
	array.data = insertFunc(array.data, mark, val)
	n = vCount
	return
}

// IndexValue 根据索引查找元素
func (array *ArrayList) IndexValue(index int) (data interface{}) {
	array.mu.Lock()
	if n := len(array.data); 0 <= index && index <= n-1 {
		data = array.data[index]
	}
	array.mu.Unlock()
	return
}

// Index 找出数据val的索引， 没找到返回-1
func (array *ArrayList) Index(val interface{}) (i int) {
	array.mu.Lock()
	defer array.mu.Unlock()

	i = -1
	if len(array.data) == 0 {
		return
	}
	for idx, ele := range array.data {
		if ele == val {
			i = idx
			break
		}
	}
	return
}

// Count 获取data的数量
func (array *ArrayList) Count(data interface{}) (n int) {
	array.mu.Lock()
	for _, ele := range array.data {
		if data == ele {
			n += 1
		}
	}
	array.mu.Unlock()
	return
}

// Remove 删除数据
// count为删除个数
// count == 0 从队首开始删除所有的data
// count > 0 从队首开始删除count个data
// count < 0 从队尾开始删除count个data
func (array *ArrayList) Remove(data interface{}, count int) (rmCount int) {
	var removeFunc = func(s []interface{}, index int) []interface{} {
		n := len(s)
		copy(s[index:], s[index+1:])
		s[n-1] = nil
		s = s[:n-1]
		return s
	}

	array.mu.Lock()
	defer array.mu.Unlock()

	switch {
	case count > 0:
		for i := 1; i <= count; i++ {
			for index, ele := range array.data {
				if ele == data {
					array.data = removeFunc(array.data, index)
					rmCount += 1
					break
				}
			}
		}
	case count < 0:
		for i := count; i < 0; i++ {
			for index, ele := range array.data {
				if ele == data {
					array.data = removeFunc(array.data, index)
					rmCount += 1
					break
				}
			}
		}
	default:
	outLoop:
		for {
			for index, ele := range array.data {
				if ele == data {
					array.data = removeFunc(array.data, index)
					rmCount += 1
					continue outLoop
				}
			}
			break outLoop
		}
	}
	return
}

// Set 设置指定位置index[0, len)的数据为val
func (array *ArrayList) Set(index int, val interface{}) (ok bool) {
	array.mu.Lock()
	defer array.mu.Unlock()

	n := len(array.data)
	if index < 0 || index >= n {
		return
	}
	array.data[index] = val
	ok = true

	return
}

// RangeFunc 遍历
func (array *ArrayList) RangeFunc(fn func(i int, val interface{}) (over bool)) {
	if fn == nil {
		return
	}
	array.mu.Lock()
	defer array.mu.Unlock()

	for i, v := range array.data {
		if fn(i, v) {
			break
		}
	}
}

// Trim 截断位置[start, end)之间的数据
func (array *ArrayList) Trim(start, end int) (result []interface{}) {
	array.mu.Lock()
	defer array.mu.Unlock()

	length := len(array.data)
	start, end = handleIndex(length, start, end)
	//start小于等于左边界，end大于等于右边界，不处理
	if start <= 0 && end >= length-1 {
		return
	}

	//start大于end，或者start超出右边界，则直接将列表置空
	if start > end || start >= length {
		result = array.data
		array.data = nil
		return
	}

	result = append(result, array.data[start:end])
	copy(array.data[start:], array.data[end:])
	for k, n := len(array.data)-end+start, len(array.data); k < n; k++ {
		array.data[k] = nil
	}
	array.data = array.data[:len(array.data)-end+start]
	return
}

func (array *ArrayList) Len() (n int) {
	array.mu.Lock()
	n = len(array.data)
	array.mu.Unlock()

	return
}

package listpkg

import (
	"container/list"
	"sync"
)

type HashList struct {
	mu   *sync.RWMutex
	data map[interface{}]*list.List
}

func NewHashList() *HashList {
	return &HashList{
		mu:   &sync.RWMutex{},
		data: make(map[interface{}]*list.List),
	}
}

func (hl *HashList) Data() (data map[interface{}][]interface{}) {
	data = make(map[interface{}][]interface{})
	for k, record := range hl.data {
		if record.Len() == 0 {
			continue
		}
		for e := record.Front(); e != nil; e = e.Next() {
			data[k] = append(data[k], e.Value)
		}
	}
	return
}

// LPush 向队首添加数据, 返回实际添加的数据个数
func (hl *HashList) LPush(key interface{}, data ...interface{}) (n int) {
	if len(data) == 0 {
		return
	}
	hl.mu.Lock()
	n = hl.push(true, key, data...)
	hl.mu.Unlock()
	return
}

// LInsert 在指定位置前面插入数据，返回实际插入的数据个数
func (hl *HashList) LInsert(key interface{}, mark interface{}, val ...interface{}) (n int) {
	if len(val) == 0 {
		return
	}
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	pElement := find(record, mark)
	if pElement == nil {
		return
	}
	for _, v := range val {
		n += 1
		record.InsertBefore(v, pElement)
	}
	return
}

// LPop 从队首取出数据
func (hl *HashList) LPop(key interface{}, n int) (elements []interface{}) {
	if n <= 0 {
		return
	}
	hl.mu.Lock()
	elements = hl.pop(true, n, key)
	hl.mu.Unlock()
	return
}

// RPush 向队尾添加数据
func (hl *HashList) RPush(key interface{}, data ...interface{}) (n int) {
	hl.mu.Lock()
	n = hl.push(false, key, data...)
	hl.mu.Unlock()
	return
}

// RPop 从队尾取数据
func (hl *HashList) RPop(key interface{}, n int) (elements []interface{}) {
	hl.mu.Lock()
	elements = hl.pop(false, n, key)
	hl.mu.Unlock()
	return
}

// RInsert 在指定位置后面插入数据
// 插入失败返回-1
// 插入成功返回队列长度
func (hl *HashList) RInsert(key interface{}, mark interface{}, val ...interface{}) (n int) {
	if len(val) == 0 {
		return
	}
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	pElement := find(record, mark)
	if pElement == nil {
		return
	}
	for _, v := range val {
		record.InsertAfter(v, pElement)
		n += 1
	}
	return
}

// IndexValue 根据索引查找数据
func (hl *HashList) IndexValue(key interface{}, idx int) (data interface{}) {
	hl.mu.RLock()
	defer hl.mu.RUnlock()

	record := hl.data[key]
	ok, newIndex := validIndex(record, idx)
	if !ok {
		return
	}

	if element := index(record, newIndex); element != nil {
		data = element.Value
	}
	return
}

// Count 获取data的数量
func (hl *HashList) Count(key interface{}, data interface{}) (n int) {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return
	}
	for e := record.Front(); e != nil; e = e.Next() {
		if e.Value == data {
			n += 1
		}
	}
	return
}

// Remove 删除数据
// count为删除个数
// count == 0 从队首开始删除所有的data
// count > 0 从队首开始删除count个data
// count < 0 从队尾开始删除count个data
func (hl *HashList) Remove(key interface{}, data interface{}, count int) (rmCount int) {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return
	}
	var es []*list.Element
	switch {
	case count > 0:
		for e := record.Front(); e != nil && len(es) < count; e = e.Next() {
			if e.Value == data {
				es = append(es, e)
			}
		}
	case count < 0:
		for e := record.Back(); e != nil && len(es) < count; e = e.Prev() {
			if e.Value == data {
				es = append(es, e)
			}
		}
	default:
		for e := record.Front(); e != nil; e = e.Next() {
			if e.Value == data {
				es = append(es, e)
			}
		}
	}
	for _, e := range es {
		record.Remove(e)
	}
	rmCount = len(es)
	return
}

// Set 将索引处的数据设置为val
func (hl *HashList) Set(key interface{}, idx int, val interface{}) (ok bool) {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	data := index(record, idx)
	if data != nil {
		data.Value = val
		ok = true
	}
	return
}

// RangeFunc 遍历start到end之间的数据
func (hl *HashList) RangeFunc(key interface{}, start, end int, fn func(data interface{}) (over bool)) {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return
	}
	length := record.Len()
	start, end = handleIndex(length, start, end)
	if start > end || start >= length {
		return
	}

	mid := length >> 1
	if end <= mid || end-mid < mid-start {
		flag := 0
		for p := record.Front(); p != nil && flag <= end; p, flag = p.Next(), flag+1 {
			if flag >= start && fn != nil {
				if fn(p.Value) {
					break
				}
			}
		}
	} else {
		flag := length - 1
		for p := record.Back(); p != nil && flag >= start; p, flag = p.Prev(), flag-1 {
			if flag <= end && fn != nil {
				if fn(p.Value) {
					break
				}
			}
		}
	}
}

// Trim 截断
func (hl *HashList) Trim(key interface{}, start, end int) bool {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	if record == nil || record.Len() <= 0 {
		return false
	}

	length := record.Len()
	start, end = handleIndex(length, start, end)

	//start小于等于左边界，end大于等于右边界，不处理
	if start <= 0 && end >= length-1 {
		return false
	}

	//start大于end，或者start超出右边界，则直接将列表置空
	if start > end || start >= length {
		hl.data[key] = nil
		delete(hl.data, key)
		return true
	}

	startEle, endEle := index(record, start), index(record, end)
	switch end-start+1 < (length >> 1) {
	case true:
		newList := list.New()
		for p := startEle; p != endEle.Next(); p = p.Next() {
			newList.PushBack(p.Value)
		}
		hl.data[key] = newList
	default:
		var ele []*list.Element
		for p := record.Front(); p != startEle; p = p.Next() {
			ele = append(ele, p)
		}
		for p := record.Back(); p != endEle; p = p.Prev() {
			ele = append(ele, p)
		}
		for _, e := range ele {
			record.Remove(e)
		}
	}
	return true
}

func (hl *HashList) Len(key interface{}) (length int) {
	hl.mu.RLock()
	length = hl.data[key].Len()
	hl.mu.RUnlock()
	return
}

// 向队列中添加元素，返回实际添加的个数
func (hl *HashList) push(front bool, key interface{}, val ...interface{}) (n int) {
	record := hl.data[key]
	if record == nil {
		record = list.New()
		hl.data[key] = record
	}

	switch {
	case front == true:
		for _, v := range val {
			record.PushFront(v)
			n++
		}
	default:
		for _, v := range val {
			record.PushBack(v)
			n++
		}
	}
	return
}

func (hl *HashList) pop(front bool, n int, key interface{}) (elements []interface{}) {
	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}
	switch {
	case front:
		for e := record.Front(); e != nil && len(elements) < n; e = e.Next() {
			elements = append(elements, e.Value)
			record.Remove(e)
		}
	default:
		for e := record.Back(); e != nil && len(elements) < n; e = e.Prev() {
			elements = append(elements, e.Value)
			record.Remove(e)
		}
	}
	return
}

func handleIndex(length, start, end int) (int, int) {
	if start < 0 {
		start += length
	}
	if end < 0 {
		end += length
	}
	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length - 1
	}
	return start, end
}

func validIndex(dataList *list.List, index int) (bool, int) {
	if dataList == nil || dataList.Len() <= 0 {
		return false, index
	}
	n := dataList.Len()
	if index < 0 {
		index += n
	}
	return index >= 0 && index < n, index
}

func find(dataList *list.List, target interface{}) (result *list.Element) {
	if dataList == nil || dataList.Len() == 0 {
		return
	}
	for ele := dataList.Front(); ele != nil; ele = ele.Next() {
		if ele.Value == target {
			result = ele
			break
		}
	}
	return
}

func index(dataList *list.List, index int) (data *list.Element) {
	if dataList == nil || dataList.Len() == 0 {
		return
	}
	ok, newIndex := validIndex(dataList, index)
	if !ok {
		return
	}
	index = newIndex

	var element *list.Element
	// 如果index在前半段，则从头开始找，否则从后半段查找
	switch index <= ((dataList.Len()) >> 1) {
	case true:
		val := dataList.Front()
		for i := 0; i < index; i++ {
			val = val.Next()
		}
		element = val
	default:
		val := dataList.Back()
		for i := dataList.Len() - 1; i > index; i-- {
			val = val.Prev()
		}
		element = val
	}
	return element
}

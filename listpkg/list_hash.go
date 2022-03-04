package listpkg

import (
	"container/list"
	"reflect"
	"sync"
)

type HashList struct {
	mu   *sync.RWMutex
	data map[interface{}]*list.List
}

func NewHList() *HashList {
	return &HashList{
		mu:   &sync.RWMutex{},
		data: make(map[interface{}]*list.List),
	}
}

// LPush 向队首添加数据
func (hl *HashList) LPush(key interface{}, data ...interface{}) (n int) {
	hl.mu.Lock()
	n = hl.push(true, key, data...)
	hl.mu.Unlock()

	return
}

// LPop 从队首取出数据
func (hl *HashList) LPop(key interface{}) (data interface{}) {
	hl.mu.Lock()
	data = hl.pop(true, key)
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
func (hl *HashList) RPop(key interface{}) (data interface{}) {
	hl.mu.Lock()
	data = hl.pop(false, key)
	hl.mu.Unlock()
	return
}

// Index 根据索引查找数据
func (hl *HashList) Index(key interface{}, idx int) (data interface{}) {
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
			if reflect.DeepEqual(e.Value, data) {
				es = append(es, e)
			}
		}
	case count < 0:
		for e := record.Back(); e != nil && len(es) < count; e = e.Prev() {
			if reflect.DeepEqual(e.Value, data) {
				es = append(es, e)
			}
		}
	default:
		for e := record.Front(); e != nil; e = e.Next() {
			if reflect.DeepEqual(e.Value, data) {
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

func (hl *HashList) UnsafeLInsert(key interface{}, pivot, val interface{}) int {
	record := hl.data[key]
	pEle := find(record, pivot)
	if pEle == nil {
		return -1
	}
	record.InsertBefore(val, pEle)
	return record.Len()
}

// LInsert 在指定位置前面插入数据
// 插入失败返回-1
// 插入成功返回队列长度
func (hl *HashList) LInsert(key interface{}, mark, val interface{}) int {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	pElement := find(record, mark)
	if pElement == nil {
		return -1
	}
	record.InsertBefore(val, pElement)
	return record.Len()
}

// RInsert 在指定位置后面插入数据
// 插入失败返回-1
// 插入成功返回队列长度
func (hl *HashList) RInsert(key interface{}, mark, val interface{}) int {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	pElement := find(record, mark)
	if pElement == nil {
		return -1
	}
	record.InsertAfter(val, pElement)

	return record.Len()
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

func (hl *HashList) UnsafeRange(key interface{}, start, end int) (result []interface{}) {
	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}
	length := record.Len()
	start, end = handleIndex(length, start, end)
	if start > end || start >= length {
		return nil
	}

	mid := length >> 1
	if end <= mid || end-mid < mid-start {
		flag := 0
		for p := record.Front(); p != nil && flag <= end; p, flag = p.Next(), flag+1 {
			if flag >= start {
				result = append(result, p.Value)
			}
		}
	} else {
		flag := length - 1
		for p := record.Back(); p != nil && flag >= start; p, flag = p.Prev(), flag-1 {
			if flag <= end {
				result = append(result, p.Value)
			}
		}
		if len(result) > 0 {
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return
}

// Range 遍历start到end之间的数据
func (hl *HashList) Range(key interface{}, start, end int) (result []interface{}) {
	hl.mu.Lock()
	defer hl.mu.Unlock()

	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}
	length := record.Len()
	start, end = handleIndex(length, start, end)
	if start > end || start >= length {
		return nil
	}

	mid := length >> 1
	if end <= mid || end-mid < mid-start {
		flag := 0
		for p := record.Front(); p != nil && flag <= end; p, flag = p.Next(), flag+1 {
			if flag >= start {
				result = append(result, p.Value)
			}
		}
	} else {
		flag := length - 1
		for p := record.Back(); p != nil && flag >= start; p, flag = p.Prev(), flag-1 {
			if flag <= end {
				result = append(result, p.Value)
			}
		}
		if len(result) > 0 {
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return
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

func (hl *HashList) push(front bool, key interface{}, val ...interface{}) int {
	record := hl.data[key]
	if record == nil {
		record = list.New()
		hl.data[key] = record
	}

	switch {
	case front == true:
		for _, v := range val {
			record.PushFront(v)
		}
	default:
		for _, v := range val {
			record.PushBack(v)
		}
	}
	return record.Len()
}

func (hl *HashList) pop(front bool, key interface{}) interface{} {
	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}
	var element *list.Element
	switch {
	case front:
		element = record.Front()
	default:
		element = record.Back()
	}
	record.Remove(element)
	return element.Value
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
		if reflect.DeepEqual(ele.Value, target) {
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

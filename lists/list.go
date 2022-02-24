package list

import (
	"container/list"
	"reflect"
	"sync"
)

type HList struct {
	mu   sync.RWMutex
	data map[interface{}]*list.List
}

func NewHList() *HList {
	return &HList{
		data: make(map[interface{}]*list.List),
	}
}

func (hl *HList) UnsafeLPush(key interface{}, data ...interface{}) int {
	hl.init()
	return hl.push(true, key, data...)
}

func (hl *HList) LPush(key interface{}, data ...interface{}) int {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeLPush(key, data...)
}

func (hl *HList) UnsafeLPop(key interface{}) interface{} {
	hl.init()
	return hl.pop(true, key)
}

func (hl *HList) LPop(key interface{}) interface{} {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeLPop(key)
}

func (hl *HList) UnsafeRPush(key interface{}, data ...interface{}) int {
	hl.init()
	return hl.push(false, key, data...)
}

func (hl *HList) RPush(key interface{}, data ...interface{}) int {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeRPush(key, data...)
}

func (hl *HList) UnsafeRPop(key interface{}) interface{} {
	hl.init()
	return hl.pop(false, key)
}

func (hl *HList) RPop(key interface{}) interface{} {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeRPop(key)
}

func (hl *HList) UnsafeIndex(key interface{}, index int) interface{} {
	hl.init()
	ok, newIndex := hl.validIndex(key, index)
	if !ok {
		return nil
	}
	index = newIndex
	element := hl.index(key, index)
	if element != nil {
		return element.Value
	}
	return nil
}

func (hl *HList) Index(key interface{}, index int) interface{} {
	hl.mu.RLock()
	defer hl.mu.RUnlock()
	return hl.UnsafeIndex(key, index)
}

func (hl *HList) UnsafeRem(key interface{}, data interface{}, count int) (rmCount int) {
	hl.init()
	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return
	}
	var es []*list.Element
	if count == 0 {
		for e := record.Front(); e != nil; e = e.Next() {
			if reflect.DeepEqual(e.Value, data) {
				es = append(es, e)
			}
		}
	}
	if count > 0 {
		for e := record.Front(); e != nil && len(es) < count; e = e.Next() {
			if reflect.DeepEqual(e.Value, data) {
				es = append(es, e)
			}
		}
	}

	if count < 0 {
		for e := record.Back(); e != nil && len(es) < count; e = e.Prev() {
			if reflect.DeepEqual(e.Value, data) {
				es = append(es, e)
			}
		}
	}

	for _, e := range es {
		record.Remove(e)
	}
	rmCount = len(es)
	es = nil

	return
}

func (hl *HList) Rem(key interface{}, data interface{}, count int) (rmCount int) {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeRem(key, data, count)
}

func (hl *HList) UnsafeLInsert(key interface{}, pivot, val interface{}) int {
	hl.init()
	pEle := hl.find(key, pivot)
	if pEle == nil {
		return -1
	}

	record := hl.data[key]
	if record == nil {
		record = list.New()
		hl.data[key] = record
	}
	record.InsertBefore(val, pEle)
	return record.Len()
}

func (hl *HList) LInsert(key interface{}, pivot, val interface{}) int {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeLInsert(key, pivot, val)
}

func (hl *HList) UnsafeRInsert(key interface{}, pivot, val interface{}) int {
	hl.init()
	pEle := hl.find(key, pivot)
	if pEle == nil {
		return -1
	}

	record := hl.data[key]
	if record == nil {
		record = list.New()
		hl.data[key] = record
	}
	record.InsertAfter(val, pEle)
	return record.Len()
}

func (hl *HList) RInsert(key interface{}, pivot, val interface{}) int {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeRInsert(key, pivot, val)
}

func (hl *HList) UnsafeSet(key interface{}, index int, val interface{}) bool {
	hl.init()
	element := hl.index(key, index)
	if element == nil {
		return false
	}
	element.Value = val
	return true
}

func (hl *HList) Set(key interface{}, index int, val interface{}) bool {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeSet(key, index, val)
}

func (hl *HList) UnsafeRange(key interface{}, start, end int) (result []interface{}) {
	hl.init()
	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}
	length := record.Len()
	start, end = hl.handleIndex(length, start, end)
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

func (hl *HList) Range(key interface{}, start, end int) (result []interface{}) {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeRange(key, start, end)
}

func (hl *HList) UnsafeTrim(key interface{}, start, end int) bool {
	hl.init()
	item := hl.data[key]
	if item == nil || item.Len() <= 0 {
		return false
	}

	length := item.Len()
	start, end = hl.handleIndex(length, start, end)

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

	startEle, endEle := hl.index(key, start), hl.index(key, end)
	if end-start+1 < (length >> 1) {
		newList := list.New()
		for p := startEle; p != endEle.Next(); p = p.Next() {
			newList.PushBack(p.Value)
		}

		item = nil
		hl.data[key] = newList
	} else {
		var ele []*list.Element
		for p := item.Front(); p != startEle; p = p.Next() {
			ele = append(ele, p)
		}
		for p := item.Back(); p != endEle; p = p.Prev() {
			ele = append(ele, p)
		}

		for _, e := range ele {
			item.Remove(e)
		}

		ele = nil
	}

	return true
}

func (hl *HList) Trim(key interface{}, start, end int) bool {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	return hl.UnsafeTrim(key, start, end)
}

func (hl *HList) UnsafeLen(key interface{}) int {
	hl.init()
	record := hl.data[key]
	if record != nil {
		return record.Len()
	}

	return 0
}

func (hl *HList) Len(key interface{}) int {
	hl.mu.RLock()
	defer hl.mu.RUnlock()
	return hl.UnsafeLen(key)
}

func (hl *HList) init() {
	if hl.data == nil {
		hl.data = make(map[interface{}]*list.List)
	}
}

func (hl *HList) push(front bool, key interface{}, val ...interface{}) int {
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

func (hl *HList) pop(front bool, key interface{}) interface{} {
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

func (hl *HList) find(key interface{}, data interface{}) *list.Element {
	record := hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}
	var e *list.Element
	for ele := record.Front(); ele != nil; ele = ele.Next() {
		if reflect.DeepEqual(ele.Value, data) {
			e = ele
			break
		}
	}
	return e
}

func (hl *HList) index(key interface{}, index int) *list.Element {
	ok, newIndex := hl.validIndex(key, index)
	if !ok {
		return nil
	}
	index = newIndex
	var record = hl.data[key]
	if record == nil || record.Len() == 0 {
		return nil
	}

	var element *list.Element
	// 如果index在前半段，则从头开始找，否则从后半段查找
	if index <= ((record.Len()) >> 1) {
		val := record.Front()
		for i := 0; i < index; i++ {
			val = val.Next()
		}
		element = val
	} else {
		val := record.Back()
		for i := record.Len() - 1; i > index; i-- {
			val = val.Prev()
		}
		element = val
	}
	return element
}

func (hl *HList) validIndex(key interface{}, index int) (bool, int) {
	record := hl.data[key]
	if record == nil || record.Len() <= 0 {
		return false, index
	}
	n := record.Len()
	if index < 0 {
		index += n
	}
	return index >= 0 && index < n, index
}

func (hl *HList) handleIndex(length, start, end int) (int, int) {
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

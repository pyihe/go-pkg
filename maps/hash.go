package maps

import "sync"

type (
	// fields key, value形式，一对一
	fields map[interface{}]interface{}

	// kv key，value形式，一对多，一个key对应多个键值对
	kv map[interface{}]fields

	// HashTable 不带锁的Hash表
	HashTable struct {
		mu     sync.RWMutex
		values kv
	}
)

func NewHashTable() *HashTable {
	return &HashTable{
		values: make(map[interface{}]fields),
	}
}

// init 初始化values
func (h *HashTable) init() {
	if h.values == nil {
		h.values = make(kv)
	}
}

// existKey 判断key是否存在
func (h *HashTable) existKey(key interface{}) bool {
	_, exist := h.values[key]
	return exist
}

// UnsafeHSet 不加锁的Set
func (h *HashTable) UnsafeHSet(key, field, value interface{}) int {
	h.init()

	if !h.existKey(key) {
		h.values[key] = make(fields)
	}
	h.values[key][field] = value
	return len(h.values[key])
}

// HSet 将key下的field字段设置为value，如果key下已经存在field字段，直接覆盖旧值，返回key下的field数量
func (h *HashTable) HSet(key, field, value interface{}) int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.UnsafeHSet(key, field, value)
}

func (h *HashTable) UnsafeHSetNX(key, field, value interface{}) (ok bool) {
	h.init()

	if !h.existKey(key) {
		h.values[key] = make(fields)
	}
	if _, exist := h.values[key][field]; !exist {
		h.values[key][field] = value
		ok = true
	}
	return false
}

// HSetNX 与HSet功能相同，不过只有当key下不存在field字段时才执行set操作
func (h *HashTable) HSetNX(key, field, value interface{}) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.UnsafeHSetNX(key, field, value)
}

func (h *HashTable) UnsafeHGet(key, field interface{}) interface{} {
	if len(h.values) == 0 {
		return nil
	} else {
		if !h.existKey(key) {
			return nil
		}
		return h.values[key][field]
	}
}

// HGet 返回key下field字段对应的值
func (h *HashTable) HGet(key, field interface{}) interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHGet(key, field)
}

func (h *HashTable) UnsafeHDel(key, field interface{}) bool {
	if len(h.values) == 0 {
		return false
	}
	if !h.existKey(key) {
		return false
	}
	if _, ok := h.values[key][field]; ok {
		delete(h.values[key], field)
		return true
	}
	return false
}

// HDel 删除key下指定的field字段，返回field是否存在
func (h *HashTable) HDel(key, field interface{}) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.UnsafeHDel(key, field)
}

func (h *HashTable) UnsafeHExists(key, field interface{}) (exist bool) {
	if len(h.values) == 0 {
		return false
	}

	if !h.existKey(key) {
		return false
	}
	_, exist = h.values[key][field]
	return
}

// HExists 判断是key是否存在field字段
func (h *HashTable) HExists(key, field interface{}) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHExists(key, field)
}

func (h *HashTable) UnsafeHLen(key interface{}) int {
	if len(h.values) == 0 {
		return 0
	} else {
		return len(h.values[key])
	}
}

// HLen 返回key对应的字段数量
func (h *HashTable) HLen(key interface{}) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHLen(key)
}

func (h *HashTable) UnsafeHGetAll(key interface{}) (values []interface{}) {
	if len(h.values) == 0 {
		return
	}
	if !h.existKey(key) {
		return
	}
	for field, value := range h.values[key] {
		values = append(values, field, value)
	}
	return
}

// HGetAll 获取key下所有的field-value值对
func (h *HashTable) HGetAll(key interface{}) (values []interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHGetAll(key)
}

func (h *HashTable) UnsafeHKeys() (keys []interface{}) {
	if len(h.values) == 0 {
		return
	}
	for k := range h.values {
		keys = append(keys, k)
	}
	return
}

// HKeys 返回所有的key
func (h *HashTable) HKeys() (keys []interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHKeys()
}

func (h *HashTable) UnsafeHFields(key interface{}) (fields []interface{}) {
	if len(h.values) == 0 {
		return
	}
	if !h.existKey(key) {
		return
	}
	for field := range h.values[key] {
		fields = append(fields, field)
	}
	return
}

func (h *HashTable) HFields(key interface{}) (fields []interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHFields(key)
}

func (h *HashTable) UnsafeHValues(key interface{}) (values []interface{}) {
	if len(h.values) == 0 {
		return
	}
	if !h.existKey(key) {
		return
	}
	for _, v := range h.values[key] {
		values = append(values, v)
	}
	return
}

// HValues 返回key包含的所有字段对应的值
func (h *HashTable) HValues(key interface{}) (values []interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.UnsafeHValues(key)
}

func (h *HashTable) UnsafeRangeKeys(fn func(key, field, value interface{}) bool) {
	if len(h.values) == 0 {
		return
	}
	for key, fvs := range h.values {
		for field, value := range fvs {
			if fn(key, field, value) {
				return
			}
		}
	}
}

func (h *HashTable) RLockRangeKeys(fn func(key, field, value interface{}) bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	h.UnsafeRangeKeys(fn)
}

func (h *HashTable) LockRangeKeys(fn func(key, field, value interface{}) bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.UnsafeRangeKeys(fn)
}

func (h *HashTable) UnsafeRangeFields(key interface{}, fn func(field, value interface{}) bool) {
	if len(h.values) == 0 {
		return
	}
	if !h.existKey(key) {
		return
	}
	for field, value := range h.values[key] {
		if fn(field, value) {
			return
		}
	}
}

func (h *HashTable) LockRangeFields(key interface{}, fn func(field, value interface{}) bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.UnsafeRangeFields(key, fn)
}

func (h *HashTable) RLockRangeFields(key interface{}, fn func(field, value interface{}) bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	h.UnsafeRangeFields(key, fn)
}

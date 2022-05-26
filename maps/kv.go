package maps

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pyihe/go-pkg/bytes"
)

type Param map[string]interface{}

func NewParam() Param {
	return make(Param)
}

func (p Param) Set(key string, value interface{}) {
	p[key] = value
}

func (p Param) SetNX(key string, value interface{}) {
	if _, exist := p[key]; !exist {
		p[key] = value
	}
}

func (p Param) MSet(ps map[string]interface{}) {
	for k, v := range ps {
		p[k] = v
	}
}

func (p Param) Del(key string) {
	delete(p, key)
}

func (p Param) Exists(key string) (ok bool) {
	_, ok = p[key]
	return
}

func (p Param) Get(key string) (value interface{}, ok bool) {
	value, ok = p[key]
	return
}

func (p Param) MGet(keys ...string) []interface{} {
	result := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		v := p[k]
		result = append(result, v)
	}
	return result
}

func (p Param) Range(fn func(key string, value interface{}) (breakOut bool)) {
	for k, v := range p {
		if fn(k, v) {
			break
		}
	}
}

func (p Param) GetString(key string) (s string, ok bool) {
	value, ok := p.Get(key)
	if !ok {
		return
	}
	switch v := value.(type) {
	case string:
		s = v
	case bool:
		if v {
			s = "1"
		} else {
			s = "0"
		}
	case []byte:
		s = bytes.String(v)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		s = strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		s = v.Format(time.RFC3339)
	case nil:
		s = ""
	default:
		s = fmt.Sprint(v)
	}
	return
}

func (p Param) GetInt64(key string) (n int64, ok bool) {
	value, ok := p.Get(key)
	if !ok {
		return
	}
	switch v := value.(type) {
	case string:
		n, _ = strconv.ParseInt(v, 10, 64)
	case bool:
		if v {
			n = 1
		} else {
			n = 0
		}
	case []byte:
		n, _ = bytes.Int64(v)
	case uint8:
		n = int64(v)
	case uint16:
		n = int64(v)
	case uint:
		n = int64(v)
	case uint32:
		n = int64(v)
	case uint64:
		n = int64(v)
	case int8:
		n = int64(v)
	case int16:
		n = int64(v)
	case int32:
		n = int64(v)
	case int64:
		n = v
	case float32:
		n = int64(v)
	case float64:
		n = int64(v)
	case time.Time:
		n = v.Unix()
	case nil:
		n = 0
	}
	return
}

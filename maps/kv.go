package maps

import (
	"reflect"
	"strconv"

	"github.com/pyihe/go-pkg/errors"
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

func (p Param) GetString(key string) (string, error) {
	value, ok := p.Get(key)
	if !ok {
		return "", errors.New("not exist key: " + key)
	}
	return reflect.ValueOf(value).String(), nil
}

func (p Param) GetInt64(key string) (n int64, err error) {
	value, ok := p.Get(key)
	if !ok {
		return 0, errors.New("not exist key: " + key)
	}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Bool:
		if v.Bool() == true {
			n = 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n = v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n = int64(v.Uint())
	case reflect.String:
		n, err = strconv.ParseInt(v.String(), 10, 64)
	case reflect.Float32, reflect.Float64:
		n = int64(v.Float())
	default:
		err = errors.New("unknown type: " + t.String())
	}
	return
}

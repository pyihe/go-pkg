package slice

import (
	"reflect"
)

type Slice interface {
	PushBack(x interface{}) (bool, int)
	PushFront(x interface{}) (bool, int)
	PopBack() (bool, interface{})
	PopFront() (bool, interface{})
	Index(x interface{}) int
	IndexValue(i int) interface{}
	Range(fn func(index int, ele interface{}) bool)
}

func NewSlice(st reflect.Kind, capacity int) Slice {
	switch st {
	case reflect.Int8:
		return newInt8Slice(capacity)
	case reflect.Uint8:
		return newUint8Slice(capacity)
	case reflect.Int16:
		return newInt16Slice(capacity)
	case reflect.Uint16:
		return newUint16Slice(capacity)
	case reflect.Int:
		return newIntSlice(capacity)
	case reflect.Uint:
		return newUintSlice(capacity)
	case reflect.Int32:
		return newInt32Slice(capacity)
	case reflect.Uint32:
		return newUint32Slice(capacity)
	case reflect.Int64:
		return newInt64Slice(capacity)
	case reflect.Uint64:
		return newUint64Slice(capacity)
	case reflect.Float32:
		return newFloat32Slice(capacity)
	case reflect.Float64:
		return newFloat64Slice(capacity)
	case reflect.String:
		return newStringSlice(capacity)
	default:
		panic("not supported type")
	}
}

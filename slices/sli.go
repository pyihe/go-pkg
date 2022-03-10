package slices

import (
	"fmt"
	"math"
	"reflect"
)

type Slice interface {
	Len() int
	Cap() int
	Sort()
	PushBack(x interface{}) (bool, int)
	PushFront(x interface{}) (bool, int)
	PopBack() (bool, interface{})
	PopFront() (bool, interface{})
	Index(x interface{}) int
	IndexValue(i int) interface{}
	Range(fn func(index int, ele interface{}) bool)
	Delete(x interface{}) (ok bool)
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

/****************************************************分界线**************************************************************/
func convertToFloat32(x interface{}) (ok bool, v float32) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = float32(d)
	case uint8:
		v = float32(d)
	case int16:
		v = float32(d)
	case uint16:
		v = float32(d)
	case int:
		v = float32(d)
	case uint:
		v = float32(d)
	case int32:
		v = float32(d)
	case uint32:
		v = float32(d)
	case int64:
		v = float32(d)
	case uint64:
		v = float32(d)
	case float32:
		v = d
	case float64:
		v = float32(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToFloat64(x interface{}) (bool, float64) {
	var v float64
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = float64(d)
	case uint8:
		v = float64(d)
	case int16:
		v = float64(d)
	case uint16:
		v = float64(d)
	case int:
		v = float64(d)
	case uint:
		v = float64(d)
	case int32:
		v = float64(d)
	case uint32:
		v = float64(d)
	case int64:
		v = float64(d)
	case uint64:
		v = float64(d)
	case float32:
		v = float64(d)
	case float64:
		v = d
	default:
		return false, 0
	}
	return true, v
}

func convertToInt(x interface{}) (bool, int) {
	var v int
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = int(d)
	case uint8:
		v = int(d)
	case int16:
		v = int(d)
	case uint16:
		v = int(d)
	case int:
		v = d
	case uint:
		v = int(d)
	case int32:
		v = int(d)
	case uint32:
		v = int(d)
	case int64:
		v = int(d)
	case uint64:
		v = int(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToUint(x interface{}) (ok bool, v uint) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		if d < 0 {
			return
		}
		v = uint(d)
	case uint8:
		v = uint(d)
	case int16:
		if d < 0 {
			return
		}
		v = uint(d)
	case uint16:
		v = uint(d)
	case int:
		if d < 0 {
			return
		}
		v = uint(d)
	case uint:
		v = d
	case int32:
		if d < 0 {
			return
		}
		v = uint(d)
	case uint32:
		v = uint(d)
	case int64:
		if d < 0 {
			return
		}
		v = uint(d)
	case uint64:
		v = uint(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToInt8(x interface{}) (ok bool, v int8) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = d
	case uint8:
		if d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case int16:
		if d < math.MinInt8 || d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case uint16:
		if d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case int:
		if d < math.MinInt8 || d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case uint:
		if d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case int32:
		if d < math.MinInt8 || d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case uint32:
		if d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case int64:
		if d < math.MinInt8 || d > math.MaxInt8 {
			return
		}
		v = int8(d)
	case uint64:
		if d > math.MaxInt8 {
			return
		}
		v = int8(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToUint8(x interface{}) (ok bool, v uint8) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		if d < 0 {
			return
		}
		v = uint8(d)
	case uint8:
		v = d
	case int16:
		if d < 0 || d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case uint16:
		if d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case int:
		if d < 0 || d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case uint:
		if d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case int32:
		if d < 0 || d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case uint32:
		if d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case int64:
		if d < 0 || d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	case uint64:
		if d > math.MaxUint8 {
			return
		}
		v = uint8(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToInt16(x interface{}) (ok bool, v int16) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = int16(d)
	case uint8:
		v = int16(d)
	case int16:
		v = d
	case uint16:
		if v > math.MaxInt16 {
			return
		}
		v = int16(d)
	case int:
		if d < math.MinInt16 || d > math.MaxInt16 {
			return
		}
		v = int16(d)
	case uint:
		if d > math.MaxInt16 {
			return
		}
		v = int16(d)
	case int32:
		if d < math.MinInt16 || d > math.MaxInt16 {
			return
		}
		v = int16(d)
	case uint32:
		if d > math.MaxInt16 {
			return
		}
		v = int16(d)
	case int64:
		if d < math.MinInt16 || d > math.MaxInt16 {
			return
		}
		v = int16(d)
	case uint64:
		if d > math.MaxInt16 {
			return
		}
		v = int16(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToUint16(x interface{}) (ok bool, v uint16) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		if d < 0 {
			return
		}
		v = uint16(d)
	case uint8:
		v = uint16(d)
	case int16:
		if d < 0 {
			return
		}
		v = uint16(d)
	case uint16:
		v = d
	case int:
		if d < 0 || d > math.MaxUint16 {
			return
		}
		v = uint16(d)
	case uint:
		if v > math.MaxUint16 {
			return
		}
		v = uint16(d)
	case int32:
		if d < 0 || d > math.MaxUint16 {
			return
		}
		v = uint16(d)
	case uint32:
		if d > math.MaxUint16 {
			return
		}
		v = uint16(d)
	case int64:
		if d < 0 || d > math.MaxUint16 {
			return
		}
		v = uint16(d)
	case uint64:
		if d > math.MaxUint16 {
			return
		}
		v = uint16(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToInt32(x interface{}) (ok bool, v int32) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = int32(d)
	case uint8:
		v = int32(d)
	case int16:
		v = int32(d)
	case uint16:
		v = int32(d)
	case int:
		v = int32(d)
	case uint:
		v = int32(d)
	case int32:
		v = d
	case uint32:
		if d > math.MaxInt32 {
			return
		}
		v = int32(d)
	case int64:
		if d < math.MinInt32 || d > math.MaxInt32 {
			return
		}
		v = int32(d)
	case uint64:
		if d > math.MaxInt32 {
			return
		}
		v = int32(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToUint32(x interface{}) (ok bool, v uint32) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		if d < 0 {
			return
		}
		v = uint32(d)
	case uint8:
		v = uint32(d)
	case int16:
		if d < 0 {
			return
		}
		v = uint32(d)
	case uint16:
		v = uint32(d)
	case int:
		if d < 0 {
			return
		}
		v = uint32(d)
	case uint:
		v = uint32(d)
	case int32:
		if d < 0 {
			return
		}
		v = uint32(d)
	case uint32:
		v = d
	case int64:
		if d < 0 || d > math.MaxUint32 {
			return
		}
		v = uint32(d)
	case uint64:
		if d > math.MaxInt32 {
			return
		}
		v = uint32(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToInt64(x interface{}) (ok bool, v int64) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		v = int64(d)
	case uint8:
		v = int64(d)
	case int16:
		v = int64(d)
	case uint16:
		v = int64(d)
	case int:
		v = int64(d)
	case uint:
		v = int64(d)
	case int32:
		v = int64(d)
	case uint32:
		v = int64(d)
	case int64:
		v = d
	case uint64:
		if d > math.MaxInt64 {
			return
		}
		v = int64(d)
	default:
		return false, 0
	}
	return true, v
}

func convertToUint64(x interface{}) (ok bool, v uint64) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = 1
		} else {
			v = 0
		}
	case int8:
		if d < 0 {
			return
		}
		v = uint64(d)
	case uint8:
		v = uint64(d)
	case int16:
		if d < 0 {
			return
		}
		v = uint64(d)
	case uint16:
		v = uint64(d)
	case int:
		if d < 0 {
			return
		}
		v = uint64(d)
	case uint:
		v = uint64(d)
	case int32:
		if d < 0 {
			return
		}
		v = uint64(d)
	case uint32:
		v = uint64(d)
	case int64:
		if d < 0 {
			return
		}
		v = uint64(d)
	case uint64:
		v = d
	default:
		return false, 0
	}
	return true, v
}

func convertToIntString(x interface{}) (ok bool, v string) {
	switch d := x.(type) {
	case bool:
		if d == true {
			v = fmt.Sprintf("1")
		} else {
			v = fmt.Sprintf("0")
		}
	case int8:
		v = fmt.Sprintf("%d", d)
	case uint8:
		v = fmt.Sprintf("%d", d)
	case int16:
		v = fmt.Sprintf("%d", d)
	case uint16:
		v = fmt.Sprintf("%d", d)
	case int:
		v = fmt.Sprintf("%d", d)
	case uint:
		v = fmt.Sprintf("%d", d)
	case int32:
		v = fmt.Sprintf("%d", d)
	case uint32:
		v = fmt.Sprintf("%d", d)
	case int64:
		v = fmt.Sprintf("%d", d)
	case uint64:
		v = fmt.Sprintf("%d", d)
	case float32:
		v = fmt.Sprintf("%f", d)
	case float64:
		v = fmt.Sprintf("%f", d)
	default:
		return false, ""
	}
	return true, v
}

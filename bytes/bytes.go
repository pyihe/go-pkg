package bytes

import (
	"fmt"
	"unsafe"

	"github.com/pyihe/go-pkg/errors"
	"github.com/valyala/bytebufferpool"
)

type ByteBuffer = bytebufferpool.ByteBuffer

var (
	Get = bytebufferpool.Get
	Put = func(b *ByteBuffer) {
		if b != nil {
			bytebufferpool.Put(b)
		}
	}
)

func Int64(b []byte) (v int64, err error) {
	if len(b) == 0 {
		return
	}
	negative := false
	if b[0] == '-' || b[0] == '+' {
		if b[0] == '-' {
			negative = true
		}
		b = b[1:]
	}
	if len(b) == 0 {
		err = errors.New("no convertable byte")
		return
	}

	for _, e := range b {
		if e < '0' || e > '9' {
			err = errors.New(fmt.Sprintf("illegal byte: %v", e))
			return
		}
		v *= 10
		v += int64(e - '0')
	}
	if negative {
		v = -v
	}
	return
}

// String []byte转换为string
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Equal 判断两个字节切片每个元素是否相等
func Equal(a, b []byte) bool {
	aLen, bLen := len(a), len(b)
	if aLen != bLen {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:aLen]
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// Contain 判断字节切片b是否包含ele元素
func Contain(b []byte, ele byte) bool {
	for _, v := range b {
		if v == ele {
			return true
		}
	}
	return false
}

// Reverse 翻转字节切片
func Reverse(b []byte) {
	l := len(b)
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - i - 1
		b[i], b[opp] = b[opp], b[i]
	}
}

// Remove 对于eles的每个元素，只删除一次
func Remove(b *[]byte, eles ...byte) {
	if b == nil || len(*b) == 0 {
		return
	}
	for _, e := range eles {
		for i := 0; i < len(*b); {
			if e == (*b)[i] {
				copy((*b)[i:], (*b)[i+1:])
				(*b)[len(*b)-1] = 0
				*b = (*b)[:len(*b)-1]
				break
			} else {
				i++
			}
		}
	}
}

// RemoveAll 删除所有的ele
func RemoveAll(b *[]byte, eles ...byte) {
	for _, e := range eles {
		for i := 0; i < len(*b); {
			if (*b)[i] == e {
				*b = append((*b)[:i], (*b)[i+1:]...)
			} else {
				i++
			}
		}
	}
}

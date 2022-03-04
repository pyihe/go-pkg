package bytepkg

import (
	"unsafe"

	"github.com/valyala/bytebufferpool"
)

type ByteBuffer = bytebufferpool.ByteBuffer

var (
	Get = bytebufferpool.Get()
	Put = func(b *ByteBuffer) {
		if b != nil {
			bytebufferpool.Put(b)
		}
	}
)

// String []byte转换为string
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// BytesEqual 判断两个字节切片每个元素是否相等
func BytesEqual(a, b []byte) bool {
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

// Remove 从b中删除ele
func Remove(b []byte, ele byte) {
	for i := range b {
		if ele == b[i] {
			b = append(b[:i], b[i+1:]...)
			break
		}
	}
}

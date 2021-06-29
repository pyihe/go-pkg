package bytes

import "unsafe"

// String []byte转换为string
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// SliceEqual
func SliceEqual(a, b []byte) bool {
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

// Contain
func Contain(ele byte, b []byte) bool {
	for _, v := range b {
		if v == ele {
			return true
		}
	}
	return false
}

// Reverse
func Reverse(b []byte) {
	l := len(b)
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - i - 1
		b[i], b[opp] = b[opp], b[i]
	}
}

// Remove
func Remove(ele byte, b []byte) {
	for i := range b {
		if ele == b[i] {
			b = append(b[:i], b[i+1:]...)
			break
		}
	}
}

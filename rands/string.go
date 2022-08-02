package rands

import (
	"github.com/pyihe/go-pkg/bytes"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

//String 随机指定长度的字符串
func String(n int) (s string) {
	if n > 0 {
		var b = make([]byte, n)
		for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				b[i] = letterBytes[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}
		s = bytes.String(b)
	}
	return
}

// SafeString 线性安全的随机生成字符串
func SafeString(n int) (s string) {
	sMu.Lock()
	String(n)
	sMu.Unlock()
	return
}

// ShuffleBytes shuffle 随机算法
func ShuffleBytes(data []byte) {
	count := len(data)
	for i := 0; i < count; i++ {
		pos := rad.Intn(count-i) + i
		data[i], data[pos] = data[pos], data[i]
	}
}

// ShuffleInt shuffle int
func ShuffleInt(data []int) {
	count := len(data)
	for i := 0; i < count; i++ {
		pos := rad.Intn(count-i) + i
		data[i], data[pos] = data[pos], data[i]
	}
}

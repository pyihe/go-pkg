package rands

import (
	"github.com/pyihe/go-pkg/bytes"
	"math/rand"
	"time"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var (
	src rand.Source
)

func init() {
	rand.Seed(time.Now().UnixNano())
	src = rand.NewSource(time.Now().UnixNano())
}

//生成min-max之间的一个随机数
func Int(min, max int) int {
	if min >= max {
		panic("min bigger than max")
	}
	return rand.Intn(max-min+1) + min
}

func Int64(min, max int64) int64 {
	if min >= max {
		panic("min bigger than max")
	}
	return rand.Int63n(max-min+1) + min
}

//随机指定长度的字符串
func String(n int) string {
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
	return bytes.String(b)
}

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
	rad *rand.Rand
)

func init() {
	src = rand.NewSource(time.Now().UnixNano())
	rad = rand.New(src)
}

// Int 生成min-max之间的一个随机数
func Int(min, max int) int {
	if min > max {
		panic("min bigger than max")
	}
	if min == max {
		return max
	}
	return rad.Intn(max-min+1) + min
}

// Int64 在min和max之间随机返回一个数
func Int64(min, max int64) int64 {
	if min > max {
		panic("min bigger than max")
	}
	if min == max {
		return min
	}
	return rad.Int63n(max-min+1) + min
}

// Int32 在min和max之间随机返回一个数
func Int32(min, max int32) int32 {
	if min > max {
		panic("min bigger than max")
	}
	if min == max {
		return min
	}
	return rad.Int31n(max-min+1) + min
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

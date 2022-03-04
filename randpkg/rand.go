package randpkg

import (
	"math/rand"
	"time"

	"github.com/pyihe/go-pkg/bytepkg"
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

//String 随机指定长度的字符串
func String(n int) string {
	if n <= 0 {
		panic("invalid argument to String")
	}
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
	return bytepkg.String(b)
}

// ShuffleBytes shuffle 随机算法
func ShuffleBytes(data []byte) {
	count := len(data)
	for i := 0; i < count; i++ {
		pos := rand.Intn(count-i) + i
		data[i], data[pos] = data[pos], data[i]
	}
}

func ShuffleInt(data []int) {
	count := len(data)
	for i := 0; i < count; i++ {
		pos := rand.Intn(count-i) + i
		data[i], data[pos] = data[pos], data[i]
	}
}

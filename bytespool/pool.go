package bytespool

import (
	"math/bits"
	"sync"
)

// copy from github.com/bytedance/gopkg/tree/develop/lang/mcache

const maxSize = 46

var caches [maxSize]sync.Pool

func init() {
	for i := 0; i < maxSize; i++ {
		size := 1 << i
		caches[i].New = func() any {
			return make([]byte, 0, size)
		}
	}
}

func Get(length int, capacity ...int) []byte {
	n := len(capacity)
	if n > 1 {
		panic("too many arguments")
	}
	c := length
	if n > 0 && capacity[0] > length {
		c = capacity[0]
	}
	buf := caches[locatePool(c)].Get().([]byte)
	buf = buf[:length]
	return buf
}

func Put(buf []byte) {
	size := cap(buf)
	if !validSize(size) {
		return
	}
	buf = buf[:0]
	caches[bsr(size)].Put(buf)
}

func locatePool(size int) int {
	if size <= 0 {
		return 0
	}
	if validSize(size) {
		return bsr(size)
	}
	return bsr(size) + 1
}

func bsr(x int) int {
	return bits.Len(uint(x)) - 1
}

func validSize(size int) bool {
	return (size & (-size)) == size
}

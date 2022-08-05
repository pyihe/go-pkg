package buffers

import (
	"bytes"
	"sync"
)

//type ByteBuffer = bytebufferpool.ByteBuffer
//
//var (
//	Get = bytebufferpool.Get
//	Put = func(b *ByteBuffer) {
//		if b != nil {
//			bytebufferpool.Put(b)
//		}
//	}
//)

var bp sync.Pool

func Get() *bytes.Buffer {
	buffer, ok := bp.Get().(*bytes.Buffer)
	if ok {
		return buffer
	}
	return &bytes.Buffer{}
}

func Put(b *bytes.Buffer) {
	if b == nil {
		return
	}
	b.Reset()
	bp.Put(b)
}

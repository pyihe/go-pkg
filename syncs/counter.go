package syncs

import (
	"sync/atomic"
)

type (
	AtomicInt32 int32
	AtomicInt64 int64
)

func (ai *AtomicInt32) Inc(delta int32) {
	atomic.AddInt32((*int32)(ai), delta)
}

func (ai *AtomicInt32) Value() int32 {
	return atomic.LoadInt32((*int32)(ai))
}

func (ai *AtomicInt64) Inc(delta int64) {
	atomic.AddInt64((*int64)(ai), delta)
}

func (ai *AtomicInt64) Value() int64 {
	return atomic.LoadInt64((*int64)(ai))
}

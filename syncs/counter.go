package syncs

import (
	"sync/atomic"
)

type (
	AtomicInt32 int32
)

func (ai *AtomicInt32) Inc(delta int32) {
	atomic.AddInt32((*int32)(ai), delta)
}

func (ai *AtomicInt32) Value() int32 {
	return atomic.LoadInt32((*int32)(ai))
}

package syncs

import "sync/atomic"

type AtomicInt int32

func (ai *AtomicInt) Inc() {
	atomic.AddInt32((*int32)(ai), 1)
}

func (ai *AtomicInt) Dec() {
	atomic.AddInt32((*int32)(ai), -1)
}

func (ai *AtomicInt) Value() int {
	return int(atomic.LoadInt32((*int32)(ai)))
}

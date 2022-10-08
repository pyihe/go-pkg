package concurrent

import (
	"sync"

	"github.com/pyihe/go-pkg/maths"
)

type Limiter struct {
	queue  chan struct{}
	waiter *sync.WaitGroup
}

func NewLimiter(size int) *Limiter {
	size = maths.MaxInt(0, size)
	return &Limiter{
		queue:  make(chan struct{}, size),
		waiter: &sync.WaitGroup{},
	}
}

func (lim *Limiter) Add(delta int) {
	switch {
	case delta > 0:
		for i := 0; i < delta; i++ {
			lim.queue <- struct{}{}
		}
	default:
		for i := 0; i > delta; i-- {
			<-lim.queue
		}
	}

	lim.waiter.Add(delta)
}

func (lim *Limiter) Done() {
	<-lim.queue
	lim.waiter.Done()
}

func (lim *Limiter) Wait() {
	lim.waiter.Wait()
}

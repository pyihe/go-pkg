package syncs

import (
	"context"
	"sync"
)

type WgWrapper struct {
	sync.WaitGroup
}

func (w *WgWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

func (w *WgWrapper) WrapWithBlock(ctx context.Context, fn func() chan error) {
	w.Add(1)
	go func(cancelCtx context.Context, errCh chan error) {
		select {
		case <-cancelCtx.Done():
			break
		case <-errCh:
			break
		}
		w.Done()
	}(ctx, fn())
}

package times

import (
	"context"
	"time"
)

// SleepWithContext 携带context的Sleep，可中断
func SleepWithContext(ctx context.Context, dur time.Duration) error {
	timer := time.NewTimer(dur)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

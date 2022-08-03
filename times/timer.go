package times

import (
	"context"
	"time"
)

type Timer interface {
	Remain(unit string) int32
	Stop()
	StopTicker()
	After(t time.Duration, handler func())
	Every(t time.Duration, handle func())
}

type myTimer struct {
	timer   *time.Timer
	ticker  *time.Ticker
	endTime time.Time
	handler func()
}

func NewTimer() Timer {
	return new(myTimer)
}

func (m *myTimer) Remain(unit string) int32 {
	if m.timer == nil {
		return 0
	}
	r := m.endTime.Sub(time.Now())
	switch unit {
	case "H":
		return int32(r.Hours())
	case "M":
		return int32(r.Minutes())
	default:
		return int32(r.Seconds())
	}
}

func (m *myTimer) Stop() {
	if m.timer == nil {
		return
	}
	m.timer.Stop()
	m.timer = nil
}

func (m *myTimer) StopTicker() {
	if m.ticker == nil {
		return
	}
	m.ticker.Stop()
	m.ticker = nil
}

func (m *myTimer) After(t time.Duration, handler func()) {
	m.endTime = time.Now().Add(t)
	m.handler = handler
	m.timer = time.AfterFunc(t, handler)
}

func (m *myTimer) Every(t time.Duration, handle func()) {
	m.ticker = time.NewTicker(t)
	go func() {
		for range m.ticker.C {
			handle()
		}
	}()
}

func Sleep(ctx context.Context, dur time.Duration) error {
	timer := time.NewTimer(dur)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

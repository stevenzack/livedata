package livedata

import (
	"context"
	"time"
)

type Time struct {
	liveData *LiveData
}

func NewTime(t time.Time) *Time {
	return &Time{
		liveData: NewLiveData(t),
	}
}

func (l *Time) Observe(lifecycleGoroutine func(), onChange func(time.Time)) {
	go l.liveData.Observe(lifecycleGoroutine, func(v interface{}) {
		onChange(v.(time.Time))
	})
}

func (l *Time) ObserveCtx(ctx context.Context, onChange func(time.Time)) {
	go l.liveData.ObserveCtx(ctx, func(v interface{}) {
		onChange(v.(time.Time))
	})
}

func (t *Time) ObserveForever(onChange func(t time.Time)) {
	t.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(time.Time))
	})
}

func (t *Time) Post(v time.Time) {
	t.liveData.Post(v)
}

func (t *Time) Get() time.Time {
	return t.liveData.Get().(time.Time)
}

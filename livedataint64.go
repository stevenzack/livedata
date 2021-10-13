package livedata

import "context"

type Int64 struct {
	liveData *LiveData
}

func NewInt64(i int64) *Int64 {
	return &Int64{
		liveData: NewLiveData(i),
	}
}

func (l *Int64) Observe(lifecycleGoroutine func(), onChange func(int64)) {
	go l.liveData.Observe(lifecycleGoroutine, func(v interface{}) {
		onChange(v.(int64))
	})
}

func (l *Int64) ObserveCtx(ctx context.Context, onChange func(int64)) {
	go l.liveData.ObserveCtx(ctx, func(v interface{}) {
		onChange(v.(int64))
	})
}

func (l *Int64) ObserveForever(onChange func(int64)) {
	go l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(int64))
	})
}

func (l *Int64) Post(i int64) {
	l.liveData.Post(i)
}

func (l *Int64) Get() int64 {
	return l.liveData.Get().(int64)
}

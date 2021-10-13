package livedata

import "context"

type Int32 struct {
	liveData *LiveData
}

func NewInt32(i int32) *Int32 {
	return &Int32{
		liveData: NewLiveData(i),
	}
}

func (l *Int32) Observe(lifecycleGoroutine func(), onChange func(int32)) {
	l.liveData.Observe(lifecycleGoroutine, func(v interface{}) {
		onChange(v.(int32))
	})
}

func (l *Int32) ObserveCtx(ctx context.Context, onChange func(int32)) {
	l.liveData.ObserveCtx(ctx, func(v interface{}) {
		onChange(v.(int32))
	})
}

func (l *Int32) ObserveForever(onChange func(int32)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(int32))
	})
}

func (l *Int32) Post(i int32) {
	l.liveData.Post(i)
}

func (l *Int32) Get() int32 {
	return l.liveData.Get().(int32)
}

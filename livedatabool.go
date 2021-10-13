package livedata

import "context"

type Bool struct {
	liveData *LiveData
}

func NewBool(b bool) *Bool {
	return &Bool{
		liveData: NewLiveData(b),
	}
}

func (l *Bool) Observe(lifecycleGoroutine func(), onChange func(bool)) {
	go l.liveData.Observe(lifecycleGoroutine, func(v interface{}) {
		onChange(v.(bool))
	})
}

func (l *Bool) ObserveCtx(ctx context.Context, onChange func(bool)) {
	go l.liveData.ObserveCtx(ctx, func(v interface{}) {
		onChange(v.(bool))
	})
}

func (l *Bool) ObserveForever(onChange func(bool)) {
	go l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(bool))
	})
}

func (l *Bool) Post(b bool) {
	l.liveData.Post(b)
}

func (l *Bool) Get() bool {
	return l.liveData.Get().(bool)
}

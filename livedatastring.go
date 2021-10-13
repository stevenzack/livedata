package livedata

import "context"

type String struct {
	liveData *LiveData
}

func NewString(s string) *String {
	return &String{
		liveData: NewLiveData(s),
	}
}

func (l *String) Observe(lifecycleGoroutine func(), onChange func(string)) {
	go l.liveData.Observe(lifecycleGoroutine, func(v interface{}) {
		onChange(v.(string))
	})
}

func (l *String) ObserveCtx(ctx context.Context, onChange func(string)) {
	go l.liveData.ObserveCtx(ctx, func(v interface{}) {
		onChange(v.(string))
	})
}

func (l *String) ObserveForever(onChange func(string)) {
	go l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(string))
	})
}

func (l *String) Post(s string) {
	l.liveData.Post(s)
}

func (l *String) Get() string {
	return l.liveData.Get().(string)
}

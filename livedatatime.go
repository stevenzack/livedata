package livedata

import "time"

type Time struct {
	l *LiveData
}

func NewTime(t time.Time) *Time {
	return &Time{
		l: NewLiveData(t),
	}
}

func (t *Time) ObserveForever(onChange func(t time.Time)) {
	t.l.ObserveForever(func(v interface{}) {
		onChange(v.(time.Time))
	})
}

func (t *Time) Post(v time.Time) {
	t.l.Post(v)
}

func (t *Time) Get() time.Time {
	return t.l.Get().(time.Time)
}

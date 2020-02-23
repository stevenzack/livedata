package livedata

type Bool struct {
	liveData *LiveData
}

func NewBool(b bool) *Bool {
	return &Bool{
		liveData: NewLiveData(b),
	}
}

func (l *Bool) ObserveForever(onChange func(bool)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(bool))
	})
}

func (l *Bool) Post(b bool) {
	l.liveData.Post(b)
}

func (l *Bool) Get() bool {
	return l.liveData.Get().(bool)
}
